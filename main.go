package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"github.com/y0c/festa-notify/handler"
)

func main() {
	_ = godotenv.Load()
	lambda.Start(handler.SendMailHandler)
}
