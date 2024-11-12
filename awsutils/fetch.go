package awsutils

import (
    "fmt"
    "log"
    "context"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/ec2"
)

func FetchResources(account string) {
    cfg, err := config.LoadDefaultConfig(context.TODO(),
        config.WithSharedConfigProfile(account),
    )
    if err != nil {
        log.Fatalf("unable to load SDK config, %v", err)
    }

    ec2Client := ec2.NewFromConfig(cfg)
    result, err := ec2Client.DescribeInstances(context.TODO(), nil)
    if err != nil {
        log.Fatalf("failed to describe instances, %v", err)
    }

    for _, reservation := range result.Reservations {
        for _, instance := range reservation.Instances {
            fmt.Printf("Instance ID: %s\n", *instance.InstanceId)
        }
    }
}
