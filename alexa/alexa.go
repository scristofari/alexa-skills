package alexa

func Speak(message string) (*AlexaResponse, error) {
	res := AlexaResponse{
		Version: "1.0",
		Response: &Response{
			OutputSpeech: &OutputSpeech{
				Type: "PlainText",
				Text: message,
			},
		},
	}
	return &res, nil
}
