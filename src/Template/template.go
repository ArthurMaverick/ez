package template

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ArthurMaverick/ez/src/util"
)

// var CFN_ENDPOINT = "https://raw.githubusercontent.com/ArthurMaverick/ez/main/Infra/CFN"
// var TF_ENDPOINT = "https://raw.githubusercontent.com/ArthurMaverick/ez/main/Infra/TF"

type Template struct {
	Resource string
	// CfnEndoint string
	// TfEndpoint string
}

func (t *Template) getEndpoints() []string {
	CfnEndoint := "https://raw.githubusercontent.com/ArthurMaverick/ez/main/Infra/CFN"
	TfEndpoint := "https://raw.githubusercontent.com/ArthurMaverick/ez/main/Infra/TF"
	endpoint := []string{CfnEndoint, TfEndpoint}
	return endpoint
}

func (t *Template) GenerateClouformationTemplates(Resource string) (err error) {
	util.CreateDir("./cloudformation")

	fullURLFile := fmt.Sprintf("%v/%v/%v.yaml", t.getEndpoints()[0], Resource, Resource)
	fileURLResponse, err := http.Get(fullURLFile)

	if err != nil {
		log.Fatalln("Download ERROR: ", err)
	}

	defer fileURLResponse.Body.Close()

	if fileURLResponse.StatusCode != http.StatusOK {
		return fmt.Errorf("bad Status: %s", fileURLResponse.Status)
	}

	dest := fmt.Sprintf("./cloudformation/%v.yaml", Resource)
	out, err := os.Create(dest)

	if err != nil {
		return err
	}

	_, err = io.Copy(out, fileURLResponse.Body)

	if err != nil {
		return err
	}

	return nil
}

func (t *Template) GenerateTerraformModules(Resource string) (err error) {
	util.CreateDir("./terraform")
	fullURLFile := fmt.Sprintf("%v/%v/%v.yaml", t.getEndpoints()[1], Resource, Resource)
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
	fmt.Println(t.getEndpoints()[0], t.getEndpoints()[1])
}
