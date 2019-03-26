package main

import (
	"fmt"
	"math/rand"
	u "github.com/nicor88/go-getting-started/pkg/utils"
	"runtime"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(u.Add(4, 5))
	fmt.Println(u.GetFibonacciSerie(3))
	fmt.Println(u.Hi("Nicola"))
	fmt.Println(u.SumLoop(5))
	fmt.Println(runtime.GOOS)
	fmt.Println("You have",runtime.NumCPU(), "cpu")
}
