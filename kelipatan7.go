package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&number)

	if number%7 == 0 {
		fmt.Printf("%d adalah kelipatan 7.\n", number)
	} else {
		fmt.Printf("%d bukan kelipatan 7.\n", number)
	}
}
