package main

import "fmt"

func main() {
	setOne := []int{10, 2, 1, 8, 14, 94, 16, 31, 23}
	setTwo := []int{30, 54, 10, 13, 1, 0, 71, 23}

	intersection := getIntersection(setOne, setTwo)
	fmt.Println(intersection)
}

func getIntersection(setOne []int, setTwo []int) []int {
	unique := make(map[int]bool)
	var intersection []int

	//ставим каждое значение из первого массива уникальным
	for _, v := range setOne {
		unique[v] = true
	}

	//проходим по 2 массиву, если в мапе уникальности есть такой же элемент, добавляем в массив пересечений
	for _, v := range setTwo {
		if _, ok := unique[v]; ok {
			intersection = append(intersection, v)
		}
	}
	return intersection
}
