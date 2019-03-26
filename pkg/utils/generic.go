package utils

import (
	"fmt"
)

// Hi returns a friendly greeting
func Hi(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

// Add it sum 2 integers
func Add(x int, y int) int {
	return x + y
}

// Fibonacci returns the number of fibonacci sequence
func Fibonacci(number int) int {
	if number == 0 {
		return 1
	} else if number == 1 {
		return number
	} else {
		return Fibonacci(number-2) + Fibonacci(number-1)
	}
}

// GetFibonacciSerie return the full Fibonacci serie
func GetFibonacciSerie(seriesLength int) []int {
	serie := make([]int, seriesLength)
	for i := 0; i < seriesLength; i++ {
		serie[i] = Fibonacci(i)
	}
	return serie
}

func SumLoop(seriesLength int) string {
	sum := 0
	for i := 0; i < seriesLength; i++ {
		sum += i
		fmt.Println(i)
	}
	seriesLengthNmber := seriesLength -1
	return fmt.Sprintf("Sum from 0 till %d is %d", seriesLengthNmber, sum)
}
