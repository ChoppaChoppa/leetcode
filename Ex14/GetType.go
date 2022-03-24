package main

import "fmt"

func main() {
	var v interface{} = 2.14
	getType(v)
}

func getType(v interface{}) {
	switch x := v.(type) {
	case int:
		fmt.Println("int:", x)

	case float64:
		fmt.Println("float64:", x)

	default:
		fmt.Println("unknown")
	}
}
