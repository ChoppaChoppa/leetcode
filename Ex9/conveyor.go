package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	chGet := make(chan int, len(arr))
	chSend := make(chan int, len(arr))
	wg := sync.WaitGroup{}

	wg.Add(1)
	go Conveyor(chGet, chSend, &wg)
	for _, v := range arr {
		chGet <- v
	}
	//ожидаем пока конвеер не сделает свое дело
	wg.Wait()

	for v := range chSend {
		fmt.Println(v)
	}
}

func Conveyor(chGet, chSend chan int, wg *sync.WaitGroup) {
	// данные будут поступать пока количество данных в канале не станет равным 0
	for len(chGet) != 0 {
		squared := <-chGet
		squared = int(math.Pow(float64(squared), 2))
		chSend <- squared
	}

	wg.Done()
	close(chSend)
	close(chGet)
	fmt.Println("conv stoped")
}
