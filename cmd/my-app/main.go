package main

import (
	"fmt"
	"math/rand"
	utils "github.com/nicor88/go-getting-started"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(utils.Add(42, 13))
}
