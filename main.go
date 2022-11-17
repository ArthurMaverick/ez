package main

import (
	"fmt"
	"os"

	t "github.com/ArthurMaverick/ez/import/template"
)

func main() {
	t := t.Template{}
	arg := os.Args
	fmt.Println(arg[1])
	fmt.Println(t.GenerateTemplate(arg[1]))
	t.PrintResource()
}
