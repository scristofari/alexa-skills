package main

import (
	"context"
	"testing"

	"github.com/scristofari/alexa-hello/alexa"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request alexa.AlexaRequest
		expect  string
	}{
		{
			request: alexa.AlexaRequest{},
			expect:  "Hello, Lambda",
		},
	}

	for _, test := range tests {
		response, _ := handler(context.Background(), test.request)
		assert.Equal(t, test.expect, response.Response.OutputSpeech.Text)
	}
}
