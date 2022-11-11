package template

import (
	"fmt"

	"github.com/ArthurMaverick/ez/util"
)

type Template struct {
	Resource string
}

func init() {
	util.CreateDir("./aws")
	util.CreateDir("./aws/cloudformation")
}

func (t *Template) GenerateTemplate(Resource string) (string, error) {

	src := fmt.Sprintf("./source/%v/%v.yaml", Resource, Resource)
	dest := fmt.Sprintf("./aws/cloudformation/%v.yaml", Resource)

	util.CopyTemplate(src, dest)

	t.Resource = Resource
	return t.Resource, nil
}

func (t *Template) PrintResource() {
	fmt.Println(t.Resource)
}
