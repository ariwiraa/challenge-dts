package main

import "fmt"

func main() {
	var n int
	fmt.Println("Masukan angka looping")
	fmt.Scan(&n)

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%d fizzbuzz \n", i)
		} else if i%3 == 0 {
			fmt.Printf("%d fizz \n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d buzz \n", i)
		} else {
			fmt.Println(i)
		}
	}
}
