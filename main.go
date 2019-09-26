package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(name string) (string, error) {
	return fmt.Sprintf("Hello %s!", name), nil
}

func main() {
	fmt.Println("Server Running")
	lambda.Start(HandleRequest)
}
