package main

import "fmt"

func main() {
	data := []string{"cat", "cat", "dog", "cat", "tree"}
	intersection := getIntersection(data)
	fmt.Println(intersection)
}

//создаем мапу и добляем туда значения, если уже есть такой ключ значение инкрементируется
func getIntersection(data []string) map[string]int {
	intersection := make(map[string]int)

	for _, v := range data {
		intersection[v]++
	}

	return intersection
}
