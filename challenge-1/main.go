package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 21
	j := true
	float := 123.456
	r := '\u042F'

	hexa := strconv.FormatInt(int64(i), 16) // ubah ke base 16
	convertInt, _ := strconv.Atoi(hexa)     // convert ke string ke int

	fmt.Printf("%v \n", i)          // nilai i
	fmt.Printf("%T \n", i)          // tipe data nilai i
	fmt.Printf("%% \n")             // mendapatkan %
	fmt.Printf("%t \n", j)          // tipe data nilai boolean true
	fmt.Printf("%t \n", !j)         // tipe data nilai boolean false
	fmt.Printf("%c \n", r)          // nilai Я
	fmt.Printf("%d \n", i)          // nilai base 10
	fmt.Printf("%o \n", i)          // nilai base 8
	fmt.Printf("%x \n", convertInt) // nilai base 16
	fmt.Printf("%X \n", convertInt) // nilai base 16
	fmt.Printf("%U \n", r)          //unicode karakter Я
	fmt.Printf("%.6f \n", float)    // nilai float
	fmt.Printf("%.6E \n", float)    // nilai float scientific

}
