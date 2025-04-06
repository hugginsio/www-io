package main

import (
	"fmt"
	"log"

	"git.huggins.io/www/internal/layouts"
)

func main() {
	log.Println("sitegen is starting")
	fmt.Println(layouts.Default("Kyle Huggins"))
}
