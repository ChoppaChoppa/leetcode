package main

import "fmt"

// два способа реализации
func main() {
	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	sortTemperature(data)
}

// более примитивный способ
// берем остаток от деления, if != 0 в зависимости от того,
// положительное ли число или нет мы прибавляем или уменьшаем
// после чего добавляем в мапу
func sortTemp(data []float64) {
	subsets := make(map[int][]float64)

	for _, v := range data {
		subset := int(v)
		for subset%10 != 0 {
			if subset > 0 {
				subset--
				continue
			}

			subset++
		}

		subsets[subset] = append(subsets[subset], v)
	}

	fmt.Println(subsets)
}

func sortTemperature(data []float64) {
	subsets := make(map[int][]float64)
	for _, v := range data {
		//этой операцией мы получаем целое число, которое можно использовать как границу
		//допусти у нас есть 27.6 разделив на 10 и переведя его в int мы получаем 2,
		//после домнажаем его на 10 и у нас получается граница
		subset := int(v/10) * 10
		subsets[subset] = append(subsets[subset], v)
	}

	fmt.Println(subsets)
}
