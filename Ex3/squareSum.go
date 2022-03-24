package main

import (
	"fmt"
	"math"
)

var arr = []float64 {2, 4, 6, 8}

func main() {
	squareSumChan := make(chan float64)

	go square(squareSumChan)
	fmt.Println(<-squareSumChan)
}

func square(squareSumChan chan float64) {
	var squareSum float64

	for _, v := range arr {
		squareSum += math.Pow(v, 2)
	}

	squareSumChan <- squareSum
}