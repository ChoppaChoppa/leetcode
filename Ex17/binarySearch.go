package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 11, 12, 13}

	fmt.Println(Search(arr, 3, true))
}

func Search(arr []int, searchNum int, sorted bool) int {
	if !sorted {
		sortedArr := arr
		quick_sorting(sortedArr, 0, len(sortedArr)-1)
	}

	return binary(arr, 0, len(arr)-1, searchNum)
}

func binary(arr []int, start, end, num int) int {
	mid := start + ((end - start) / 2)
	fmt.Println("mid: ", mid, "start: ", start, "end: ", end)

	if arr[mid] < num {
		fmt.Println("first")
		return binary(arr, mid+1, end, num)
	} else if arr[mid] > num {
		fmt.Println("second")
		return binary(arr, start, mid-1, num)
	}

	return mid
}

func quick_sorting(data []int, start, end int) {
	if start < end {
		base := data[start]
		left := start
		right := end

		for left < right {
			for left < right && data[right] >= base {
				right--
			}

			if left < right {
				data[left] = data[right]
				left++
			}

			for left < right && data[left] <= base {
				left++
			}

			if left < right {
				data[right] = data[left]
				right--
			}
		}

		data[left] = base
		quick_sorting(data, start, left-1)
		quick_sorting(data, left+1, end)
	}
}
