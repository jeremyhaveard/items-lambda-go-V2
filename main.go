package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	//"incoding/json"
)

//type ResponseBody struct {
//	Message string `json:"message"`
//}

func main() {
	lambda.Start(handler)
	log.Println("Items API is spooling up")
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Items API was called")

	var response = events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "{name:'Item 1'}",
	}

	return response, nil

}
