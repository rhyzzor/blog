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
	title, err := extractString(metaData, "title")
	if err != nil {
		return Post{}, err
	}

	slug, err := extractString(metaData, "slug")
	if err != nil {
		return Post{}, err
	}

	date, err := extractString(metaData, "date")
	if err != nil {
		return Post{}, err
	}

	tags, err := extractTags(metaData)
	if err != nil {
		return Post{}, err
	}

	return Post{
		Title:   title,
		Slug:    slug,
		Content: template.HTML(buf.String()),
		Tags:    tags,
		Date:    date,
	}, nil
}

func extractString(metaData map[string]interface{}, key string) (string, error) {
	value, ok := metaData[key].(string)
	if !ok {
		return "", fmt.Errorf("%s must be a string", key)
	}
	return value, nil
}

func extractTags(metaData map[string]interface{}) ([]string, error) {
	tagsInterface, ok := metaData["tags"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("tags must be an array of strings")
	}

	tags := make([]string, len(tagsInterface))
	for i, tag := range tagsInterface {
		tagStr, ok := tag.(string)
		if !ok {
			return nil, fmt.Errorf("tag %d must be a string", i)
		}
		tags[i] = tagStr
	}

	return tags, nil
}
