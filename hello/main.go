package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Body string `json:"body"`
}

type Response struct {
	IsBase64Encoded bool              `json:"isBase64Encoded"`
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	Body            string            `json:"body"`
}

type Payload struct {
	Name string `json:"name"`
}

func Handler(request Request) (Response, error) {
	data := &Payload{
		Name: "",
	}
	json.Unmarshal([]byte(request.Body), data)

	messageBody := struct {
		Message string `json:"name"`
	}{
		fmt.Sprintf("Hello %v thank you for calling me!", data.Name),
	}

	result, _ := json.Marshal(messageBody)

	return Response{
		StatusCode:      200,
		Body:            string(result),
		IsBase64Encoded: false,
		Headers:         make(map[string]string),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
