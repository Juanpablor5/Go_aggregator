package kinesisproducer

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	producer "github.com/mitooos/kinesis-producer"
)

var instance *producer.Producer

func GetProducer() *producer.Producer {
	if instance == nil {
		instance = newProducer()
	}

	return instance
}

func newProducer() *producer.Producer {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-west-2"),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: os.Getenv("AWS_ACCESS_ID"), SecretAccessKey: os.Getenv("AWS_ACCESS_KEY"),
				Source: "environment variables credentials",
			},
		}),
	)
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	client := kinesis.NewFromConfig(cfg)

	pr := producer.New(&producer.Config{
		StreamName:   os.Getenv("DESTINATION_STREAM"),
		BacklogCount: 2000,
		Client:       client,
	})

	pr.Start()
	// Handle failures
	go func() {
		for r := range pr.NotifyFailures() {
			// r contains `Data`, `PartitionKey` and `Error()`
			log.Printf("detected put failure, %v", r)
		}
	}()

	return pr
}
