package main

import (
	"fmt"
	"log"
	"os"

	t "github.com/ArthurMaverick/ez/src/Template"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	access_key string = os.Getenv("AWS_ACCESS_KEY_ID")
	secret_key string = os.Getenv("AWS_SECRET_ACCESS_KEY")
)

func main() {
	_, err := session.NewSession()

	if err != nil {
		log.Fatalln(err)
	}

	if access_key == "" || secret_key == "" {
		log.Fatalln("Credentials is Empty")
	}

	t := t.Template{}
	arg := os.Args
	fmt.Println(arg[1])
	fmt.Println(t.GenerateClouformationTemplates(arg[1]))
	t.PrintResource()
}
