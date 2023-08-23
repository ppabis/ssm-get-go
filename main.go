package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <ParameterPath> <OutputFile>\n", os.Args[0])
		os.Exit(0)
	}

	region := "eu-central-1"
	if os.Getenv("AWS_REGION") != "" {
		region = os.Getenv("AWS_REGION")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		fmt.Printf("[ERROR] unable to load SDK config, %v\n", err)
	}

	parameterName := os.Args[1]
	fileName := os.Args[2]

	client := ssm.NewFromConfig(cfg)

	mTrue := true
	output, err := client.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name:           &parameterName,
		WithDecryption: &mTrue,
	})

	if err != nil {
		fmt.Printf("[ERROR] getting parameter %s: %s\n", parameterName, err)
		os.Exit(0)
	}

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("[ERROR] creating file %s: %s\n", fileName, err)
		os.Exit(0)
	}

	defer file.Close()

	_, err = file.WriteString(*output.Parameter.Value)
	if err != nil {
		fmt.Printf("[ERROR] writing file %s: %s\n", fileName, err)
		os.Exit(0)
	}
}
