package quote

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

// Quote struct
type Quote struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

// GetQuote returns a random quote
func GetQuote() Quote {

	quote := getQuoteFromDb(getRamdomID())
	return quote
}

func getRamdomID() string {
	// min can be a fixed value because we start with our ID in the database with 1
	var min int64 = 1
	// TODO: move to env variable
	var max int64 = DescribeTableByName("motivation-quote-go")
	// generate a random integer between min and max while min is the minimum value and not 0
	id := rand.Int63n(max-min+1) + min
	// convert the integer to string because DynamoDB only supports float values
	return strconv.FormatInt(id, 10)
}

func getQuoteFromDb(id string) Quote {
	log.Printf("Get quote with id: %v", id)
	items := QueryQuoteByID(id)
	quotes := []Quote{}
	unmarshalError := attributevalue.UnmarshalListOfMaps(items, &quotes)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	log.Printf("Quote: %v", quotes)
	// it should be always 1 quote
	return quotes[0]
}
