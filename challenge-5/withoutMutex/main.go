package main

import (
	"fmt"
	"sync"
)

func main() {
	bisa := []interface{}{"bisa1", "bisa2", "bisa3"}
	coba := []interface{}{"coba1", "coba2", "coba3"}

	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go printCoba(i, coba, &wg)
		go printBisa(i, bisa, &wg)
	}

	wg.Wait()
}

func printCoba(i int, coba []interface{}, wg *sync.WaitGroup) {
	fmt.Println(coba, i)

	wg.Done()
}

func printBisa(i int, bisa []interface{}, wg *sync.WaitGroup) {
	fmt.Println(bisa, i)

	wg.Done()
}
