package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request events.APIGatewayProxyRequest
		expect  string
	}{
		{
			request: events.APIGatewayProxyRequest{},
			expect:  "Hello Lambda",
		},
	}

	for _, test := range tests {
		response, _ := handler(test.request)
		assert.Equal(t, test.expect, response.Body)
	}
}
