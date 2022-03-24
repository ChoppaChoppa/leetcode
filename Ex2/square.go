package main

import (
	"fmt"
	"math"
)

var arr = [5]float64{2, 4, 6, 8, 10}

func main() {
	squareChan := make(chan [5]float64)

	go square(squareChan)

	fmt.Println(<-squareChan)
}

func square(squareChan chan [5]float64) {
	squareArr := [5]float64{}
	for i, v := range arr {
		squareArr[i] = math.Pow(v, 2)
	}

	squareChan <- squareArr
}
