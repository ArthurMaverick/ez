package cfn

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/aws-sdk-go/aws"
)

type Cloudformation struct {
}

var cfn *cloudformation.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		panic("configuration error, " + err.Error())
	}
	cfn = cloudformation.NewFromConfig(cfg)
}

func EzCreateStack(stackName string, templateBody string) {

	cfn.CreateStack(context.TODO(), &cloudformation.CreateStackInput{
		StackName:    aws.String("test"),
		TemplateBody: aws.String(templateBody),
		Parameters:   []types.Parameter{},
	})
}
