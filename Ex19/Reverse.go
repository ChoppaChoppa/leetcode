package main

import "fmt"

func main() {
	fmt.Println(reverse("Hello World abcd"))
}

func reverse(arg string) string {
	var reverseStr string

	for i := 0; i < len(arg); i++ {
		reverseStr += string(arg[len(arg)-i-1])
	}

	return reverseStr
}
