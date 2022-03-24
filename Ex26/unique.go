package main

import "fmt"

func main() {
	arg := "avifksutopWQeczxDAsdwqqwdag"

	fmt.Println(CheckUniqueArr(arg))

	check, first, second := CheckUniqueMap(arg)
	fmt.Println("check: ", check, "; first: ", first, "; second: ", second)
}

// чере массив.каждый символ проверяется с прошлыми
func CheckUniqueArr(arg string) bool {
	var checkedArg []int32
	for _, vArg := range arg {
		if len(checkedArg) == 0 {
			checkedArg = append(checkedArg, vArg)
			continue
		}

		for _, v := range checkedArg {
			if v == vArg || v+32 == vArg {
				return false
			}
		}
	}

	return true
}

// через мпау. добавляется если в мапе нет такого ключа, иначе false
func CheckUniqueMap(arg string) (bool, int, int) {
	checkedChar := make(map[int32]int)
	for i, v := range arg {
		iUpper, okUpper := checkedChar[v]
		if okUpper {
			return false, iUpper, i
		}

		iLower, okLower := checkedChar[v+32]
		if okLower {
			return false, iLower, i
		}

		checkedChar[v] = i
	}

	return true, 0, 0
}
