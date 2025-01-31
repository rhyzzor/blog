package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/gin-gonic/gin"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

type Post struct {
	Title       string
	Slug        string
	Description string
	Content     template.HTML
	Date        time.Time
	URL         string
}

const url = "https://rhyzzor.com/"

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"transformToShort": transformToShort,
		"transformToLong":  transformToLong,
		"setTitle":         setTitle,
		"calculateReadingTime": func(content template.HTML) int {
			return calculateReadingTime(content)
		},
	})

	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	posts, err := loadMarkdown("markdown")
	if err != nil {
		log.Fatalf("failed to load markdown: %v", err)
	}

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Posts":       posts,
			"Title":       "",
			"URL":         url,
			"Description": "A blog to document my journey",
		})
	})

	r.GET("/books", func(c *gin.Context) {
		c.HTML(http.StatusOK, "books.html", gin.H{
			"Title":       "Books",
			"URL":         url + "books",
			"Description": "This is a list of book that I've read",
		})
	})

	for _, post := range posts {
		r.GET(fmt.Sprintf("/%s", post.Slug), func(c *gin.Context) {
			c.HTML(http.StatusOK, "post.html", gin.H{
				"Post": post,
			})
		})
	}

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	r.Run()
}

func sortByDate(posts []Post) []Post {
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})

	return posts
}

func loadMarkdown(dir string) ([]Post, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var posts []Post
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".md" {
			continue
		}

		wg.Add(1)
		go func(file os.DirEntry) {
			defer wg.Done()

			content, err := os.ReadFile(filepath.Join(dir, file.Name()))
			if err != nil {
				log.Printf("failed to read file %s: %v", file.Name(), err)
				return
			}

			post, err := parseMarkdownToHTML(content)
			if err != nil {
				log.Printf("failed to parse markdown in file %s: %v", file.Name(), err)
				return
			}

			mu.Lock()
			posts = append(posts, post)
			mu.Unlock()
		}(file)
	}

	wg.Wait()
	return sortByDate(posts), nil
}

func parseMarkdownToHTML(content []byte) (Post, error) {
	var buf bytes.Buffer
	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
			highlighting.NewHighlighting(
				highlighting.WithStyle("onedark"),
				highlighting.WithGuessLanguage(true),
				highlighting.WithFormatOptions(
					chromahtml.WithLineNumbers(true),
				),
			),
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)

	context := parser.NewContext()
	if err := markdown.Convert(content, &buf, parser.WithContext(context)); err != nil {
		return Post{}, fmt.Errorf("failed to convert markdown: %w", err)
	}

	metaData := meta.Get(context)
	title, err := extractString(metaData, "title")
	if err != nil {
		return Post{}, err
	}

	description, err := extractString(metaData, "description")
	if err != nil {
		return Post{}, err
	}

	slug, err := extractString(metaData, "slug")
	if err != nil {
		return Post{}, err
	}

	date, err := extractDate(metaData, "date")
	if err != nil {
		return Post{}, err
	}

	return Post{
		Title:       title,
		Slug:        slug,
		Content:     template.HTML(buf.String()),
		Description: description,
		Date:        date,
		URL:         url + slug,
	}, nil
}

func extractString(metaData map[string]interface{}, key string) (string, error) {
	value, ok := metaData[key].(string)
	if !ok {
		return "", fmt.Errorf("%s must be a string", key)
	}
	return value, nil
}

func extractDate(metaData map[string]interface{}, key string) (time.Time, error) {
	value, ok := metaData[key].(string)
	if !ok {
		return time.Time{}, fmt.Errorf("%s must be a string", key)
	}

	return time.Parse("2006-01-02 15:04:05", value)
}

func transformToShort(date time.Time) string {
	return date.Format("2006-01-02")
}

func transformToLong(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

func extractStringFromHTML(content template.HTML) string {
	p := strings.NewReader(string(content))

	doc, _ := goquery.NewDocumentFromReader(p)

	doc.Find("pre").Remove()

	return doc.Text()
}

func calculateReadingTime(content template.HTML) int {
	text := extractStringFromHTML(content)

	words := strings.Fields(text)

	return len(words) / 200
}

func setTitle(text string) string {
	if text != "" {
		return text + " | Ryan Vieira's Blog"
	}

	return "Ryan Vieira's Blog"
}
