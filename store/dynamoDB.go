package store

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/subramanian0331/urlShortener/models"
	"os"
)

const UrlTable = "UrlTable"

var ErrorNewSession = errors.New("couldn't create new dynamo db session")
var ErrorUrlNotFound = errors.New("url not found")

type dynamoDB struct {
	host    string
	port    string
	region  string
	session *dynamodb.DynamoDB
}

// NewDynamoDB constructor for db object
func NewDynamoDB(host string, port string, region string) (IStorage, error) {
	db := dynamoDB{
		host:   host,
		port:   port,
		region: region,
	}
	var err error
	db.session, err = db.NewSession()
	if err != nil {
		fmt.Println("here", err.Error())
		return nil, err
	}
	return &db, nil
}

func (d *dynamoDB) NewSession() (*dynamodb.DynamoDB, error) {
	dbHost := os.Getenv("DB_URL")
	dbPort := os.Getenv("DB_PORT")
	endPoint := fmt.Sprintf("http://%s:%s", dbHost, dbPort)
	fmt.Println(endPoint)
	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(endPoint),
		Region:   aws.String(d.region), // Replace with your desired region
	})
	if err != nil {
		return nil, ErrorNewSession
	}
	svc := dynamodb.New(sess)

	_, err = svc.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(UrlTable),
	})
	if err == nil {
		return svc, nil
	}

	createTableInput := &dynamodb.CreateTableInput{
		TableName: aws.String(UrlTable),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("shortUrl"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("shortUrl"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
	}
	_, err = svc.CreateTable(createTableInput)
	if err != nil {
		return nil, err
	}
	return svc, nil
}

func (d *dynamoDB) Save(url *models.URL) error {

	data, err := dynamodbattribute.MarshalMap(&url)
	if err != nil {
		return err
	}
	_, err = d.session.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(UrlTable),
		Item:      data,
	})
	if err != nil {
		return err
	}
	return nil
}

func (d *dynamoDB) Get(shortCode string) (*models.URL, error) {
	result, err := d.session.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(UrlTable),
		Key: map[string]*dynamodb.AttributeValue{
			"shortUrl": {S: aws.String(shortCode)},
		},
	})
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, ErrorUrlNotFound
	}

	var url models.URL
	err = dynamodbattribute.UnmarshalMap(result.Item, &url)
	if err != nil {
		return nil, err
	}
	return &url, nil
}
