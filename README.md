# motivation-quote-go

[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/lukas-weiss/motivation-quote-go/blob/master/LICENSE)

This project contains source code and supporting files for a serverless application to get a random motivation quote from famous athletes.

## Development

Following the [Google Styleguide](https://google.github.io/styleguide/go/)

Using

- gofmt
- golint

## Quotes

Quotes are stored in DynamoDB expecting a table with the name `motivation-quote-go`

| Attribute        | Name   | Type   |
| ---------------- | ------ | ------ |
| partitionKey     | id     | string |
| sortKey          | author | string |
| additional field | quote  | string |

### Example item of the DynamoDB table

| id  | author     | quote              |
| --- | ---------- | ------------------ |
| "1" | "Joe Gurt" | "nice to meet you" |

## Rquired environment variables

- AWS_REGION - the AWS region automatically available in AWS Lambda
- QUOTE_TABLE_NAME - the name of the DynamoDB table
- MAX_QUOTES - the maximum number of quotes in the database
