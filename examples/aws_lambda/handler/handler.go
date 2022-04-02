package handler

import (
	"context"
	"fmt"
	"os"

	"github.com/Code-Hex/dd"
	"github.com/aws/aws-lambda-go/lambda"
)

type Input struct {
	A string
	B int
}

func handler(ctx context.Context, input Input) {
	fmt.Println(dd.Dump(input))
}

func Run() {
	os.Setenv("_LAMBDA_SERVER_PORT", "8080")
	os.Setenv("AWS_LAMBDA_RUNTIME_API", "AWS_LAMBDA_RUNTIME_API")
	lambda.Start(handler)
}
