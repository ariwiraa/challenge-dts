package main

import (
	"fmt"
	"sync"
)

func main() {
	bisa := []interface{}{"bisa1", "bisa2", "bisa3"}
	coba := []interface{}{"coba1", "coba2", "coba3"}

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go Print(i, coba, bisa, &wg, &mutex)
	}

	wg.Wait()
}

func Print(i int, coba []interface{}, bisa []interface{}, wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	fmt.Println(coba, i)
	fmt.Println(bisa, i)
	mutex.Unlock()
	wg.Done()
}
