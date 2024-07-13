package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lukas-weiss/motivation-quote-go/internal/quote"
	"github.com/lukas-weiss/motivation-quote-go/internal/response"
)

func main() {
	lambda.Start(handler)
}

func handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("%+v\n", event)
	quote := quote.GetQuote()
	repsonse := response.CreateResponse(quote)
	return repsonse, nil
}
