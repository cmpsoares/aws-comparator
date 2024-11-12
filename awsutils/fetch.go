package awsutils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// ResourceFetcher defines a function type for fetching resources
type ResourceFetcher func(context.Context, aws.Config) ([]interface{}, error)

// FetchResources fetches resources from the specified AWS account and services.
func FetchResources(account string, services []string, outputFormat string) {
	// Load AWS Config
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(account),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Map of services to their fetchers
	resourceFetchers := map[string]ResourceFetcher{
		"ec2":            fetchEC2Instances,
		"s3":             fetchS3Buckets,
		"iam":            fetchIAMUsers,
		"rds":            fetchRDSInstances,
		"lambda":         fetchLambdaFunctions,
		"cloudformation": fetchCloudFormationStacks,
		"dynamodb":       fetchDynamoDBTables,
	}

	for _, service := range services {
		if fetcher, ok := resourceFetchers[service]; ok {
			results, err := fetcher(context.TODO(), cfg)
			if err != nil {
				log.Printf("Error fetching %s resources: %v\n", service, err)
				continue
			}

			printResources(results, service, outputFormat)
		} else {
			log.Printf("Unsupported service: %s\n", service)
		}
	}
}

// Fetch functions for individual AWS services

func fetchEC2Instances(ctx context.Context, cfg aws.Config) ([]interface{}, error) {
	client := ec2.NewFromConfig(cfg)
	output, err := client.DescribeInstances(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resources []interface{}
	for _, reservation := range output.Reservations {
		for _, instance := range reservation.Instances {
			resources = append(resources, instance)
		}
	}
	return resources, nil
}

func fetchS3Buckets(ctx context.Context, cfg aws.Config) ([]interface{}, error) {
	client := s3.NewFromConfig(cfg)
	output, err := client.ListBuckets(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resources []interface{}
	for _, bucket := range output.Buckets {
		resources = append(resources, bucket)
	}
	return resources, nil
}

func fetchIAMUsers(ctx context.Context, cfg aws.Config) ([]interface{}, error) {
	client := iam.NewFromConfig(cfg)
	output, err := client.ListUsers(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resources []interface{}
	for _, user := range output.Users {
		resources = append(resources, user)
	}
	return resources, nil
}

func fetchRDSInstances(ctx context.Context, cfg aws.Config) ([]interface{}, error) {
	client := rds.NewFromConfig(cfg)
	output, err := client.DescribeDBInstances(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resources []interface{}
	for _, dbInstance := range output.DBInstances {
		resources = append(resources, dbInstance)
	}
	return resources, nil
}

func fetchLambdaFunctions(ctx context.Context, cfg aws.Config) ([]interface{}, error) {
	client := lambda.NewFromConfig(cfg)
	output, err := client.ListFunctions(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resources []interface{}
	for _, function := range output.Functions {
		resources = append(resources, function)
	}
	return resources, nil
}

func fetchCloudFormationStacks(ctx context.Context, cfg aws.Config) ([]interface{}, error) {
	client := cloudformation.NewFromConfig(cfg)
	output, err := client.DescribeStacks(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resources []interface{}
	for _, stack := range output.Stacks {
		resources = append(resources, stack)
	}
	return resources, nil
}

func fetchDynamoDBTables(ctx context.Context, cfg aws.Config) ([]interface{}, error) {
	client := dynamodb.NewFromConfig(cfg)
	output, err := client.ListTables(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resources []interface{}
	for _, tableName := range output.TableNames {
		resources = append(resources, tableName)
	}
	return resources, nil
}

// Print resources in the desired format
func printResources(resources []interface{}, service string, format string) {
	fmt.Printf("Resources for service: %s\n", service)
	switch format {
	case "json":
		data, _ := json.MarshalIndent(resources, "", "  ")
		fmt.Println(string(data))
	default:
		for _, resource := range resources {
			fmt.Printf("%+v\n", resource)
		}
	}
}
