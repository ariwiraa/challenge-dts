package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i)

	}
	for j := 0; j <= 10; j++ {
		fmt.Println("Nilai j = ", j)
		if j == 4 {
			//САШАРВО
			var chars = []rune{'C', 'A', 'Ш', 'A', 'Р', 'В', 'О'}
			for i := 0; i < len(chars); i++ {
				char := chars[i]
				fmt.Printf("character %#U start at position byte %d\n", char, i*2)
			}
		}
	}

}
