package main

import (
	"fmt"
	"time"
)

func main() {
	gorouChanStop()
	goRouChanStopWithFeedBack()
}

// если в горутину придут данные из канала, она завершится
func gorouChanStop() {
	ch := make(chan bool)

	go func(chanal chan bool) {
		for {
			select {
			default:
				fmt.Println("goroutine work")
				time.Sleep(time.Second * 1)

			case <-chanal:
				fmt.Println("goroutine stoped")
				return
			}
		}
	}(ch)

	time.Sleep(time.Second * 5)
	ch <- false
}

// также завершается через канал, но гоурутина дает обратный фидбек, например что она завершилась
func goRouChanStopWithFeedBack() {
	ch := make(chan string)

	go func(chanal chan string) {
		for {
			select {
			case <-chanal:
				chanal <- "stoped gorou 2"
				return
			}
		}
	}(ch)

	ch <- ""

	fmt.Println(<-ch)
}

// я знаю еще про остановку через context, но в основе там лежал тот же канал, решил не использовать
// но если надо могу сразу на сдаче реализовать
