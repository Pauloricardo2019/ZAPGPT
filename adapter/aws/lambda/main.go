package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"zapgpt/adapter/aws/gateway"
)

func main() {
	lambda.Start(gateway.Process)
}
