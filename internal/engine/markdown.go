package engine

import (
	"bufio"
	"bytes"
	"errors"
	"io/fs"
	"log"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type MarkdownEngine struct {
	document    ast.Node
	extensions  parser.Extensions
	frontmatter Frontmatter
	info        fs.FileInfo
	renderer    markdown.Renderer
}

type Frontmatter struct {
	Date       string
	Layout     string
	Published  bool
	Subheading string
	Tags       []string
	Title      string
}

// The markdown engine parses and renders markdown files (to HTML by default) while extracting document frontmatter.
func Markdown() *MarkdownEngine {
	return &MarkdownEngine{
		extensions: parser.CommonExtensions | parser.AutoHeadingIDs,
		renderer:   html.NewRenderer(html.RendererOptions{Flags: html.CommonFlags}),
	}
}

func (e *MarkdownEngine) Name() string {
	return strings.Replace(e.info.Name(), filepath.Ext(e.info.Name()), ".html", 1)
}

func (e *MarkdownEngine) Parse(contents []byte, info fs.FileInfo) {
	e.info = info
	buffer := bytes.NewBuffer(contents)
	scanner := bufio.NewScanner(buffer)

	var bufferYaml bytes.Buffer
	var bufferMarkdown bytes.Buffer

	scanner.Scan()
	if strings.HasPrefix(scanner.Text(), "---") {
		inFrontmatter := true
		for scanner.Scan() {
			line := scanner.Text()
			if inFrontmatter {
				if line == "---" {
					inFrontmatter = false
					continue
				}

				bufferYaml.WriteString(line + "\n")
			} else {
				bufferMarkdown.WriteString(line + "\n")
			}
		}
	}

	parser := parser.NewWithExtensions(e.extensions)

	if bufferMarkdown.Len() == 0 {
		e.document = parser.Parse(contents)
	} else {
		e.document = parser.Parse(bufferMarkdown.Bytes())
		var frontmatter Frontmatter
		if err := yaml.Unmarshal(bufferYaml.Bytes(), &frontmatter); err != nil {
			log.Println("WARN failed to parse frontmatter:", err)
		}

		e.frontmatter = frontmatter
	}
}

func (e *MarkdownEngine) WithRenderer(renderer markdown.Renderer) {
	e.renderer = renderer
}

func (e *MarkdownEngine) Render() ([]byte, error) {
	if e.document == nil {
		return nil, errors.New("document is nil")
	}

	if e.renderer == nil {
		return nil, errors.New("renderer is nil")
	}

	// TODO: merge generated HTML partial from markdown render with page layouts
	// TODO: minify generated html

	return markdown.Render(e.document, e.renderer), nil
}
