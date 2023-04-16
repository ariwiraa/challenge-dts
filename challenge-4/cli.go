package main

import (
	"fmt"
	"os"
	"strconv"
)

type person struct {
	Name    string
	Address string
	Job     string
	Reason  string
}

func main() {

	persons := []person{
		{
			Name:    "ari",
			Address: "bekasi",
			Job:     "mahasiswa",
			Reason:  "iseng",
		},
		{
			Name:    "wira",
			Address: "jakarta",
			Job:     "freelance",
			Reason:  "iseng",
		},
	}

	args := os.Args // ambil argumen yang diinput

	convert, err := strconv.Atoi(args[1]) // konversi ke integer
	if err != nil {
		fmt.Println("Argumen bukan angka")
		return
	}

	person := persons[convert-1] // index di slice -1
	fmt.Printf("Nama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan: %s\n", person.Name, person.Address, person.Job, person.Reason)
}
