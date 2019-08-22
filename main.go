package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/y0c/festa-notify/handler"
)

func main() {
	lambda.Start(handler.SendMailHandler)
}
