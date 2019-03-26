package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var addCases = []struct {
	X int
	Y int
	Expected int
}{
	{
		X: 5,
		Y: 3,
		Expected: 8,
	},
	{
		X: 100,
		Y: 3,
		Expected: 103,
	},

}

func TestAdd(t *testing.T) {
	for _, tc := range addCases{
		result := Add(tc.X, tc.Y)
		assert.Equal(t, tc.Expected, result)
	}
}

var fibonacciCases = []struct {
	Length int
	Expected []int
}{
	{
		Length: 3,
		Expected: []int{1, 1, 2},
	},
	{
		Length: 4,
		Expected: []int{1, 1, 2, 3},
	},
	{
		Length: 5,
		Expected: []int{1, 1, 2, 3, 5},
	},
}

func TestGetFibonacciSerie(t *testing.T) {
	for _, tc := range fibonacciCases{
		result := GetFibonacciSerie(tc.Length)
		assert.Equal(t, tc.Expected, result)
	}
}

var sumLoopCases =[]struct {
	Length int
	Expected string
}{
	{
		Length: 2,
		Expected: "Sum from 0 till 1 is 1",
	},
}

func TestSumLoop(t *testing.T) {
	for _, tc := range sumLoopCases{
		result := SumLoop(tc.Length)
		assert.Equal(t, tc.Expected, result)
	}
}

var hiCases = []struct{
	Name string
	Expected string
}{
	{
		Name: "Nicola",
		Expected: "Hi, Nicola",
	},
}

func TestHi(t *testing.T) {
	for _, tc := range hiCases{
		result := Hi(tc.Name)
		assert.Equal(t, tc.Expected, result)
	}
}
