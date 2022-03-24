package main

import (
	"fmt"
	"strconv"
)

var justString []string

func someFunc() {
	v := createHugeString(1 << 7)
	justString = v[:100]

	fmt.Println(justString)
}

func main() {
	someFunc()

}

func createHugeString(b int) []string {
	var someStr []string
	for i := 0; i < b; i++ {
		someStr = append(someStr, "test"+strconv.Itoa(i))
	}

	return someStr
}
