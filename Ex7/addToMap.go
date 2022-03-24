package main

import (
	"fmt"
	"sync"
)

type CunMap struct {
	mute *sync.Mutex
	data map[int]int
}

func main() {
	mapa := CunMap{
		mute: &sync.Mutex{},
		data: make(map[int]int),
	}
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go mapa.Add(i, i+10, &wg)
	}

	wg.Wait()
	fmt.Println(mapa.data[15])
}

// добавление данных с использованием mutex
func (mapa CunMap) Add(key int, value int, wg *sync.WaitGroup) {
	mapa.mute.Lock()
	mapa.data[key] = value
	mapa.mute.Unlock()

	wg.Done()
}
