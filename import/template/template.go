package template

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ArthurMaverick/ez/util"
)

var ENDPOINT = "https://raw.githubusercontent.com/ArthurMaverick/ez/main/source"

type Template struct {
	Resource string
}

func init() {
	util.CreateDir("./aws")
	util.CreateDir("./aws/cloudformation")
}

func (t *Template) GenerateTemplate(Resource string) (err error) {
	fullURLFile := fmt.Sprintf("%v/%v/%v.yaml", ENDPOINT, Resource, Resource)

	fileURLResponse, err := http.Get(fullURLFile)

	if err != nil {
		log.Fatalln("Download ERROR: ", err)
	}

	defer fileURLResponse.Body.Close()

	if fileURLResponse.StatusCode != http.StatusOK {
		return fmt.Errorf("bad Status: %s", fileURLResponse.Status)
	}

	// src := fmt.Sprintf("./source/%v/%v.yaml", Resource, Resource)
	dest := fmt.Sprintf("./aws/cloudformation/%v.yaml", Resource)
	out, err := os.Create(dest)

	if err != nil {
		return err
	}

	_, err = io.Copy(out, fileURLResponse.Body)

	if err != nil {
		return err
	}

	// util.CopyTemplate(src, dest)

	return nil
}

func (t *Template) PrintResource() {
	fmt.Println(t.Resource)
}
