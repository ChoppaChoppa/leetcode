package main

import "fmt"

func main() {
	var num1, num2 = 20, 30
	swap(&num1, &num2)
	fmt.Println(num1, num2)
}

func swap(num1, num2 *int) (int, int) {
	*num1 += *num2
	*num2 = *num1 - *num2
	*num1 = *num1 - *num2
	return *num1, *num2
}
