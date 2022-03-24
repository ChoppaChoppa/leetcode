package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	sleep(int64(time.Second * 5))
	fmt.Println("end")
}

func sleep(duration int64) {
	t := time.Now().Add(time.Duration(duration))
	for t.After(time.Now()) {
	}
}
