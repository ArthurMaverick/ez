package ecs

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
)

type ECS struct{}

func (c *ECS) Config(sess *session.Session, region, access_key, secret_key string) *ecs.ECS {
	ecs := ecs.New(sess, &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(access_key, secret_key, ""),
	})
	return ecs
}

func (c *ECS) GetClusters(client *ecs.ECS) (*ecs.DescribeClustersOutput, error) {
	params := &ecs.DescribeClustersInput{}
	payload, error := client.DescribeClusters(params)

	if error != nil {
		log.Fatalln(error)
		return nil, error
	}
	return payload, nil
}

func (c *ECS) GetServices(client *ecs.ECS, clusterName string) (*ecs.DescribeServicesOutput, error) {
	params := &ecs.DescribeServicesInput{
		Cluster: aws.String(clusterName),
	}
	payload, error := client.DescribeServices(params)

	if error != nil {
		log.Fatalln(error)
		return nil, error
	}
	return payload, nil
}

// func (c *ECS) getServicesArn(client *ecs.ECS, clusterName string) (*ecs.ListServicesOutput, error) {
// 	params := &ecs.ListServicesInput{
// 		Cluster: aws.String(clusterName),
// 	}
// 	payload, error := client.(params)

// 	if error != nil {
// 		log.Fatalln(error)
// 		return nil, error
// 	}
// 	return payload, nil
// }

func (c *ECS) GetTasks(client *ecs.ECS, clusterName, serviceName string) (*ecs.DescribeTasksOutput, error) {
	params1 := &ecs.ListTasksInput{
		Cluster:     aws.String(clusterName),
		ServiceName: aws.String(serviceName),
	}
	payload1, error := client.ListTasks(params1)

	if error != nil {
		log.Fatalln(error)
		return nil, error
	}

	params2 := &ecs.DescribeTasksInput{
		Cluster: aws.String(clusterName),
		Tasks:   payload1.TaskArns,
	}
	payload2, error := client.DescribeTasks(params2)

	if error != nil {
		log.Fatalln(error)
		return nil, error
	}
	return payload2, nil
}

func (c *ECS) GetContainers(client *ecs.ECS, clusterName string) (*ecs.DescribeContainerInstancesOutput, error) {
	params := &ecs.DescribeContainerInstancesInput{
		Cluster: aws.String(clusterName),
	}
	payload, error := client.DescribeContainerInstances(params)

	if error != nil {
		log.Fatalln(error)
		return nil, error
	}
	return payload, nil
}

// todo
func (c *ECS) UpdateTask(client *ecs.ECS, clusterName, serviceName string) (*ecs.UpdateTaskSetOutput, error) {
	params1 := &ecs.TaskSet{
		ClusterArn: aws.String(clusterName),
		ServiceArn: aws.String(""),
	}

	params2 := &ecs.UpdateTaskSetInput{
		Cluster: aws.String(clusterName),
		Service: aws.String(serviceName),
		TaskSet: params1.TaskSetArn,
	}
	payload, error := client.UpdateTaskSet(params2)

	if error != nil {
		log.Fatalln(error)
		return nil, error
	}
	return payload, nil
}
