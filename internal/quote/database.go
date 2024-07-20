package quote

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const quoteTableNamEnv = "QUOTE_TABLE_NAME"

func getConnection() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = os.Getenv("AWS_REGION")
		return nil
	})

	if err != nil {
		// TODO: optimize error handling
		panic(err)
	}
	svc := dynamodb.NewFromConfig(cfg)
	return svc
}

// QueryQuoteByID retrieves a quote from the database based on the provided ID.
// Parameter:
// - id: the ID of the quote to retrieve.
// Returns an interface{} containing the retrieved quote.
func QueryQuoteByID(id string) []map[string]types.AttributeValue {
	svc := getConnection()
	result, err := svc.Query(context.TODO(), &dynamodb.QueryInput{
		TableName: aws.String(os.Getenv(quoteTableNamEnv)),
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

	return result.Items
}

// DescribeQuoteTable retrieves information about a DynamoDB table by name.
// Returns:
// - int64: the item count of the table
func DescribeQuoteTable() int64 {
	// infos in describe table only getting update every 6 hours
	svc := getConnection()
	result, err := svc.DescribeTable(context.TODO(), &dynamodb.DescribeTableInput{
		TableName: aws.String(os.Getenv(quoteTableNamEnv)),
	})
	if err != nil {
		panic(err)
	}

	return *result.Table.ItemCount
}
