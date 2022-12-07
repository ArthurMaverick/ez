package ssm

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

type ParameterStore struct{}

var sess *session.Session

func init() {
	payload, err := session.NewSession()

	if err != nil {
		log.Fatalln(err)
	}
	sess = payload
}

func (p *ParameterStore) Config(region, access_key, secret_key string) *ssm.SSM {
	ssm := ssm.New(sess, &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(access_key, secret_key, ""),
	})
	return ssm
}

func (p *ParameterStore) AddParameterValue(svc *ssm.SSM, name string, value string, type_value string) {
	params := &ssm.PutParameterInput{
		Name:      aws.String(name),
		Value:     aws.String(value),
		Type:      aws.String(type_value),
		Overwrite: aws.Bool(true),
	}

	req, output := svc.PutParameterRequest(params)
	err := req.Send()
	if err != nil {
		log.Fatalln("function error: ", err)
	}

	fmt.Println(req.Send())
	fmt.Println(output.String())
}

func (p *ParameterStore) GetParameterValue(svc *ssm.SSM, name string) (*string, error) {
	params := &ssm.GetParameterInput{
		Name: aws.String(name),
	}

	value, err := svc.GetParameter(params)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return value.Parameter.Value, nil
}

func (p *ParameterStore) GetParametersValues(svc *ssm.SSM, names []string) (*ssm.GetParametersOutput, error) {
	params := &ssm.GetParametersInput{
		Names: aws.StringSlice(names),
	}
	value, err := svc.GetParameters(params)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return value, nil
}

func (p *ParameterStore) DeleteParameter(svc *ssm.SSM, name string) (*ssm.DeleteParameterOutput, error) {
	params := &ssm.DeleteParameterInput{
		Name: aws.String(name),
	}
	value, err := svc.DeleteParameter(params)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return value, nil
}

func (p *ParameterStore) DeleteParameters(svc *ssm.SSM, names []string) (*ssm.DeleteParametersOutput, error) {
	params := &ssm.DeleteParametersInput{
		Names: aws.StringSlice(names),
	}
	value, err := svc.DeleteParameters(params)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return value, nil
}
