package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/xlab/closer"
)

func worker(ch chan string, wNum int) {
	var out string

	// аналог defer, призванный для кореткного закрытия worrker'а
	closer.Bind(func() {
		if out != "" {
			fmt.Println(out)
		}
		fmt.Println("worker closed ", wNum)
	})

	for {
		out = <-ch
		fmt.Println("its print worker number : ", wNum, out)
		time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	}
}

// метод который перехватывает нажатие клавишь
// при нажатии ctrl + c вызывает метод close,
// который вызывет функции добавленные в closer.Bind -
// аналог defer, но он вызывается даже если мы сами закрываем программу
func StopProgramm() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	fmt.Println("Press ESC to quit")
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)
		if key == keyboard.KeyCtrlC {
			closer.Close()
		}
	}
}

func main() {
	go StopProgramm()
	var WorkerCount int
	fmt.Scan(&WorkerCount)

	ch := make(chan string)
	for i := 0; i < WorkerCount; i++ {
		go worker(ch, i)
	}

	for {
		ch <- strconv.Itoa(rand.Intn(100))
		time.Sleep(time.Second / 2)
	}
}
