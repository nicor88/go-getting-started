package main

import (
	"fmt"
	"github.com/nicor88/go-getting-started/pkg/aws"
	"github.com/nicor88/go-getting-started/pkg/utils"
	"time"
)

type MyConfig struct {
	Environment string
	Version     string
	HostName    string
}

// this ovveride the Sting interface for the object MyConf
func (c *MyConfig) String() string {
	return fmt.Sprintf("Environment: '%v'\nVersion:'%v'\nHostName: '%v'",
		c.Environment, c.Version, c.HostName)
}


func init() {
	startingTime := utils.UnixMillisToRFC3339(time.Now().UTC().UnixNano() / int64(time.Millisecond))
	fmt.Println("Initializing the app...")
	fmt.Println("Starting at:", startingTime)
	aws.DynamoDBSession = aws.CreateDynamoDBSession()
}

type M map[interface{}]interface{}

func main() {
	ddb := &aws.DynamoDB{}
	ddb.GetClient()

	tables, _ := ddb.GetAllTables()

	for _, table := range tables {
		fmt.Println("Table name: ", *table)
	}

	tableName := *tables[0] // pick the first table

	items, _ := ddb.GetItems(tableName)

	fmt.Println("There are ", len(items), "items in table", tableName)

	cleanItems := aws.ProcessDynamoItems(items)

	fmt.Println("Processed", len(cleanItems), "items")

	// mapD := map[string]int{"apple": 5, "lettuce": 7}
	// mapB, _ := json.Marshal(mapD)
	// fmt.Println(string(mapB))

	//jsonDoc :=`
    //    {
    //        "Environment" : "Dev",
    //        "Version" : ""
    //    }`
	//
	//conf := &MyConfig{}
	//
	//json.Unmarshal([]byte(jsonDoc), conf)
	//
	//fmt.Println(conf)

	//fmt.Println("My favorite number is", rand.Intn(10))
	//fmt.Println(u.Add(4, 5))
	//fmt.Println(u.GetFibonacciSerie(3))
	//fmt.Println(u.Hi("Nicola"))
	//fmt.Println(u.SumLoop(5))
	//fmt.Println(runtime.GOOS)
	//fmt.Println("You have",runtime.NumCPU(), "cpu")
}
