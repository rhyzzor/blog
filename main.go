package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

type Post struct {
	Title   string
	Slug    string
	Content template.HTML
	Tags    []string
	Date    string
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	loadMarkdown("markdown")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run()
}

func loadMarkdown(dir string) {

	files, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	for _, file := range files {
		content, err := os.ReadFile(dir + "/" + file.Name())
		if err != nil {
			return
		}
		parseMarkdownToHTML(content)
	}
}

func parseMarkdownToHTML(content []byte) (Post, error) {
	var buf bytes.Buffer
	markdown := goldmark.New(
		goldmark.WithExtensions(meta.Meta),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)

	context := parser.NewContext()
	if err := markdown.Convert(content, &buf, parser.WithContext(context)); err != nil {
		return Post{}, fmt.Errorf("failed to convert markdown: %w", err)
	}

	metaData := meta.Get(context)
	title, ok := metaData["title"].(string)
	if !ok {
		return Post{}, fmt.Errorf("title must be a string")
	}
	slug, ok := metaData["slug"].(string)
	if !ok {
		return Post{}, fmt.Errorf("slug must be a string")
	}
	tagsInterface, ok := metaData["tags"].([]interface{})
	if !ok {
		return Post{}, fmt.Errorf("tags must be an array")
	}
	tags := make([]string, len(tagsInterface))
	for i, tag := range tagsInterface {
		tagStr, ok := tag.(string)
		if !ok {
			return Post{}, fmt.Errorf("all tags must be strings")
		}
		tags[i] = tagStr
	}
	date, ok := metaData["date"].(string)
	if !ok {
		return Post{}, fmt.Errorf("date must be a string")
	}

	return Post{
		Title:   title,
		Slug:    slug,
		Content: template.HTML(buf.String()),
		Tags:    tags,
		Date:    date,
	}, nil
}
