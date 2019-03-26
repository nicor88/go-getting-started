package main

import (
	"fmt"
	"math/rand"
	u "github.com/nicor88/go-getting-started/pkg/utils"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(u.Add(4, 5))
	fmt.Println(u.GetFibonacciSerie(10))
	fmt.Println(u.Hi("Nicola"))
}
