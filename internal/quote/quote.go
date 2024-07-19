package quote

import (
	"context"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Quote struct
type Quote struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

type record struct {
	ID     string `json:"id"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

// GetQuote returns a random quote
func GetQuote() Quote {
	// TODO: calculate the range based on the number of quotes, is there a way without a scan?
	min := 1
	max := 3
	// generate a random integer between min and max while min is the minimum value and not 0
	id := rand.Intn(max-min+1) + min
	// convert the integer to string because DynamoDB only supports float values
	quote := getQuoteFromDb(strconv.Itoa(id))
	return quote
}

func getQuoteFromDb(id string) Quote {
	log.Printf("Get quote with id: %v", id)
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = os.Getenv("AWS_REGION")

		return nil
	})
	if err != nil {
		// TODO: optimize error handling
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)
	result, err := svc.Query(context.TODO(), &dynamodb.QueryInput{
		// TODO: how to have not hardcoded value
		TableName: aws.String("motivation-quote-go"),
		KeyConditions: map[string]types.Condition{
			"id": {
				ComparisonOperator: types.ComparisonOperatorEq,
				AttributeValueList: []types.AttributeValue{
					&types.AttributeValueMemberS{Value: id},
				},
			},
		},
	})

	if err != nil {
		panic(err)
	}

	quotes := []Quote{}
	unmarshalError := attributevalue.UnmarshalListOfMaps(result.Items, &quotes)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	log.Printf("Quote: %v", quotes)
	return quotes[0]
}
