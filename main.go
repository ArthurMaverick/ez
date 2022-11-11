package main

import (
	"fmt"

	t "github.com/ArthurMaverick/ez/import/template"
)

func main() {
	t := t.Template{}

	fmt.Println(t.GenerateTemplate("vpc"))
	t.PrintResource()
}
