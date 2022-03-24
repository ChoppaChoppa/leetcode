package main

import "fmt"

func main() {
	arr := []int{6, 3, 8, 1, 0, 3, 9, 2, 5}
	quickSort(arr, 0, len(arr)-1)

	fmt.Println(arr)
}

func quickSort(arr []int, start, end int) {
	left, right := parti(arr, start, end)

	if start < right {
		quickSort(arr, start, right)
	}
	if left < end {
		quickSort(arr, left, end)
	}
}

func parti(arr []int, start, end int) (int, int) {
	mid := (start + end) / 2
	markerLeft := start
	markerRight := end
	fmt.Println("mid: ", mid)
	for i := start; markerLeft < markerRight; i++ {
		for arr[markerLeft] < mid {
			markerLeft++
		}

		for arr[markerRight] > mid {
			markerRight--
		}

		if markerLeft <= markerRight {
			temp := arr[markerLeft]
			arr[markerLeft] = arr[markerRight]
			arr[markerRight] = temp
			markerLeft++
			markerRight--
		}
	}

	return markerLeft, markerRight
}
