package main

import (
	"fmt"
	"html/template"
)

func main() {
	tmpl := template.New("tmpl")
	fmt.Print(tmpl)
}
