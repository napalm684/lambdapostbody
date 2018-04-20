package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// MessageProvider defines a contract for
// retrieving messages with text.
type MessageProvider interface {
	SetMessageText(text string) *MessageBody
}

// MessageBody defines a message structure
// containing text.
type MessageBody struct {
	Text string `json:"message"`
}

// SetMessageText implements MessageProvider by providing
// a mechanism in which one can set the message text.
func (m *MessageBody) SetMessageText(text string) *MessageBody {
	m.Text = text
	return m
}

// Request object defines the body property
// for appropriate json un-marshaling.
type Request struct {
	Body string `json:"body"`
}

// Response object defines required AWS API-Gateway
// format for getting data back to the caller.
type Response struct {
	IsBase64Encoded bool              `json:"isBase64Encoded"`
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	Body            string            `json:"body"`
}

// Handler handles api requests from an AWS API-Gateway
// and returns a message utilizing the parameter in the request body.
func Handler(request Request) (Response, error) {
	var data struct {
		Name string `json:"name"`
	}
	json.Unmarshal([]byte(request.Body), &data)

	messageBody := &MessageBody{}
	buildMessage(data.Name, messageBody)
	result, _ := json.Marshal(messageBody)

	return Response{
		StatusCode:      200,
		Body:            string(result),
		IsBase64Encoded: false,
		Headers:         make(map[string]string),
	}, nil
}

func buildMessage(p string, v MessageProvider) {
	v.SetMessageText(fmt.Sprintf("Hello %v thank you for calling me!", p))
}

func main() {
	lambda.Start(Handler)
}
