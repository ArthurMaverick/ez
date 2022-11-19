package main

import (
	"fmt"
	"os"

	t "github.com/ArthurMaverick/ez/src/template"
)

func main() {
	t := t.Template{}
	arg := os.Args
	fmt.Println(arg[1])
	fmt.Println(t.GenerateClouformationTemplates(arg[1]))
	t.PrintResource()
}
