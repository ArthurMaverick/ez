package template

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ArthurMaverick/ez/package/util"
)

type Template struct {
	Resource string
}

func (t *Template) GetEndpoints() []string {
	CfnEndoint := "https://raw.githubusercontent.com/ArthurMaverick/ez/main/package/Infra/CFN"
	TfEndpoint := "https://raw.githubusercontent.com/ArthurMaverick/ez/main/package/Infra/TF"
	endpoint := []string{CfnEndoint, TfEndpoint}
	return endpoint
}

func (t *Template) GenerateClouformationTemplates(Resource string) (err error) {
	util.CreateDir("./cloudformation")
	fullURLFile := fmt.Sprintf("%v/%v/%v.yaml", t.GetEndpoints()[0], Resource, Resource)
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
	fullURLFile := fmt.Sprintf("%v/%v", t.GetEndpoints()[1], Resource)
	fileURLResponse, err := http.Get(fullURLFile)
	fmt.Println(fileURLResponse)

	if err != nil {
		log.Fatalln("Download ERROR: ", err)
	}

	defer fileURLResponse.Body.Close()

	if fileURLResponse.StatusCode != http.StatusOK {
		return fmt.Errorf("bad Status: %s", fileURLResponse.Status)
	}

	dest := fmt.Sprintf("./terraform/%v", Resource)
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
	fmt.Println(t.GetEndpoints()[0], t.GetEndpoints()[1])
}
