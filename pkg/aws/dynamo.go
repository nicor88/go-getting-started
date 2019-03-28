package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)


var (
	DynamoDBSession *session.Session
)

func CreateDynamoDBSession() *session.Session {
	return session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
}

type DynamoDB struct {
	Client dynamodbiface.DynamoDBAPI
}


func (d *DynamoDB) GetClient() {
	d.Client = dynamodb.New(DynamoDBSession)
}


func (d *DynamoDB) GetAllTables() ([]*string, error) {
	result, err := d.Client.ListTables(&dynamodb.ListTablesInput{})
	tables := result.TableNames
	if err != nil {
		fmt.Println(err)

	}
	return tables, err
}

func (d *DynamoDB) GetItems(TableName string) ([]map[string]*dynamodb.AttributeValue, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(TableName),
	}

	result, err := d.Client.Scan(input)
	if err != nil {
		fmt.Println(err)

	}

	output := result.Items
	return output, err
}

type M map[interface{}]interface{}

func ProcessDynamoItems(items []map[string]*dynamodb.AttributeValue) []M {
	var cleanItems []M

	for _, item := range items {

		cleanItem := M {}

		for key, value :=range item {
			a := value
			//fmt.Println("Key:", key, "Value:", *a.S )

			// TODO check all the other data types
			if a.S != nil {
				cleanItem[key] = *a.S
			}

		}

		cleanItems = append(cleanItems, cleanItem)
		fmt.Println(cleanItem)
	}
	return cleanItems
}


func DynamoListTables() ([]*string, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	result, err := svc.ListTables(&dynamodb.ListTablesInput{})
	tables := result.TableNames

	if err != nil {
		fmt.Println(err)

	}
	return tables, err
}
