package main

import (
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"git.huggins.io/www/internal/engine"
)

func main() {
	log.Println("sitegen is starting")
	Preflight()

	inputPath := "content/"
	if err := filepath.Walk(inputPath, func(path string, info fs.FileInfo, err error) error {
		if path == inputPath {
			return nil
		}

		log.Println("Processing", path)

		raw, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// TODO: content engine determination should probably be in the engine package
		var ce engine.Engine
		switch filepath.Ext(info.Name()) {
		case ".md":
			ce = engine.Markdown()
		default:
			ce = engine.Default()
		}

		ce.Parse(raw, info)
		outputPath := strings.Replace(path, inputPath, "_output/", 1)
		outputPath = strings.Replace(outputPath, filepath.Base(outputPath), ce.Name(), 1)
		content, err := ce.Render()
		if err != nil {
			return err
		}

		if err := os.WriteFile(outputPath, content, 0755); err != nil {
			return err
		}

		return nil
	}); err != nil {
		log.Fatalln(err)
	}

	log.Println("thanks for flying sitegen")
}

func Preflight() {
	Delete("_output")
	Directory("_output", 0755)
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
