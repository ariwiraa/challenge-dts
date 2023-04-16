package main

import (
	"fmt"
	"strings"
)

func main() {
	sample := "selamat malam"

	results := calcWord(sample)
	fmt.Println(results)
}

func calcWord(sample string) map[string]int {
	results := make(map[string]int)
	for _, v := range sample {
		letter := string(v)

		fmt.Printf("%s\n", letter)

		count := strings.Count(sample, letter) // hitung jumlah huruf yang sama

		results[letter] = count // memasukan data ke dalam map

	}

	return results
}
