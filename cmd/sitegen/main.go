package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"

	"git.huggins.io/www/internal/layouts"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func main() {
	log.Println("sitegen is starting")

	Delete("_output")
	Directory("_output", 0755)

	// TODO: replace with pages-based markdown-driven page & path generation
	// TODO: optimize HTML & css
	pages := map[string]gomponents.Node{
		"index": layouts.Base(
			"Kyle Huggins",
			"Software engineer & Kubernetes enthusiast. Occasional photographer.",
			[]gomponents.Node{},
			[]gomponents.Node{
				html.H1(
					gomponents.Text(strings.ToUpper("huggins.io.")),
				),
				html.H2(
					gomponents.Text(strings.ToUpper("Name")),
				),
				html.H2(
					gomponents.Text(strings.ToUpper("Description")),
				),
				html.H2(
					gomponents.Text(strings.ToUpper("See Also")),
				),
			},
		),
	}

	for name, page := range pages {
		filepath := fmt.Sprintf("_output/%s.html", name)
		f := File(filepath)
		defer f.Close()

		log.Println("rendering", filepath)

		if err := page.Render(f); err != nil {
			log.Fatalln(err)
		}
	}

}

func Delete(path string) {
	if err := os.RemoveAll(path); err != nil && errors.Is(err, fs.ErrNotExist) {
		log.Fatalln(err)
	}
}

func Directory(path string, perm os.FileMode) {
	if err := os.Mkdir(path, perm); err != nil {
		log.Fatalln(err)
	}
}

func File(path string) *os.File {
	f, err := os.Create(path)

	if err != nil {
		log.Fatalln(err)
	}

	return f
}
