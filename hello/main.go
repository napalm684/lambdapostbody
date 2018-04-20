package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// Request object defines the body property
// for appropriate json un-marshaling.
type Request struct {
	Body string `json:"body"`
}

// Response object defines required AWS Api-Gateway
// format for getting object data back to the caller.
type Response struct {
	IsBase64Encoded bool              `json:"isBase64Encoded"`
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	Body            string            `json:"body"`
}

// Handler handles api requests from an AWS Api-Gateway
// and returns a message utilizing the parameter in the request body.
func Handler(request Request) (Response, error) {
	return Response{
		StatusCode:      200,
		Body:            getResponseBody(&request),
		IsBase64Encoded: false,
		Headers:         make(map[string]string),
	}, nil
}

func getName(request *Request) string {
	var data struct {
		Name string `json:"name"`
	}
	json.Unmarshal([]byte(request.Body), &data)
	return data.Name
}

func buildMessageText(p string) string {
	return fmt.Sprintf("Hello %v thank you for calling me!", p)
}

func getResponseBody(request *Request) string {
	messageBody := struct {
		Text string `json:"message"`
	}{
		buildMessageText(getName(request)),
	}
	result, _ := json.Marshal(messageBody)
	return string(result)
}

func main() {
	lambda.Start(Handler)
}
