package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/scristofari/alexa-hello/alexa"
)

func handler(ctx context.Context, r alexa.AlexaRequest) (*alexa.AlexaResponse, error) {
	return alexa.Speak("Hello, Lambda")
}
func main() {
	lambda.Start(handler)
}
