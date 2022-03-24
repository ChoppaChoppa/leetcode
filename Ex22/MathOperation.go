package main

import (
	"fmt"
	"math"
)

type command int

var (
	plus     command = 0
	minus    command = 1
	multiply command = 2
	divide   command = 3
)

func main() {
	fmt.Println(mathOperation(math.Pow(300, 29), math.Pow(200, 19), divide))
}

func mathOperation(x float64, y float64, operator command) (float64, bool) {
	if x > math.Pow(2, 20) && y > math.Pow(2, 20) {
		switch operator {
		case plus:
			return x + y, true
		case minus:
			return x - y, true
		case multiply:
			return x * y, true
		case divide:
			return x / y, true
		}
	}

	return 0, false
}
