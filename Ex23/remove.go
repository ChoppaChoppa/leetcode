package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	slice = remove(slice, 3)
	fmt.Println(slice)
}

func remove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}
