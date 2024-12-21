package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var num float64
	var numbers []float64

	for {
		_, err := fmt.Fscan(os.Stdin, &num)
		if err != nil {
			os.Exit(0)
		}
		numbers = append(numbers, num)
		min := Average(Chunk(numbers)) - StandardDeviation(Chunk(numbers))
		max := Average(Chunk(numbers)) + StandardDeviation(Chunk(numbers))
		if len(Chunk(numbers)) > 1 {
			fmt.Println(min, max)
		}

	}
}

func Average(s []float64) float64 {
	var sum float64
	for _, value := range s {
		sum += value
	}

	average := sum / float64(len(s))
	return average
}

func StandardDeviation(s []float64) float64 {
	avrg := Average(s)
	var standardDeviation float64
	var sum float64
	var variance float64

	for _, value := range s {
		sum += ((value - avrg) * (value - avrg))
		variance = sum / float64(len(s))
	}
	standardDeviation = math.Sqrt(variance)
	return standardDeviation
}

func Chunk(arr []float64) []float64 {
	ini := len(arr) - 4
	if ini < 0 {
		ini = 0
	}
	newArray := arr[ini:]
	return newArray
}
