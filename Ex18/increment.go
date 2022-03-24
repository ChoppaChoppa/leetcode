package main

import (
	"fmt"
	"sync"
)

type increment struct {
	iMutex sync.Mutex
	i      int
}

func main() {
	wg := new(sync.WaitGroup)
	i := increment{i: 0}
	for j := 0; j < 10; j++ {
		wg.Add(1)
		go i.inc(wg)
	}

	wg.Wait()
	fmt.Println("inc: ", i.i)
}

func (inc *increment) inc(wg *sync.WaitGroup) {
	inc.iMutex.Lock()
	defer inc.iMutex.Unlock()
	defer wg.Done()

	inc.i++
}
