package response

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

type ResponseBody struct {
	Message string `json:"message"`
}

func CreateResponse(quote any) events.APIGatewayProxyResponse {

	jbytes, error := json.Marshal(quote)
	jstring := string(jbytes)
	if error != nil {
		log.Printf("error: %v", error)
		return events.APIGatewayProxyResponse{StatusCode: 500}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       jstring,
	}
}
