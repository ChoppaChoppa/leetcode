package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	chan1 := make(chan int)
	_ = time.AfterFunc(time.Second*10, func() {
		os.Exit(0)
	})

	go write(chan1)
	read(chan1)
}

func write(chan1 chan int) {
	for {
		chan1 <- rand.Intn(100)
		time.Sleep(time.Second * 1)
	}

}

func read(chan2 chan int) {
	for {
		fmt.Println(<-chan2)
	}
}
