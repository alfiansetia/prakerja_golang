package main

import (
	"fmt"
)

func main() {
	var alas, tinggi, sisiMiring float64

	fmt.Print("Masukkan panjang alas: ")
	fmt.Scan(&alas)
	fmt.Print("Masukkan tinggi trapesium: ")
	fmt.Scan(&tinggi)
	fmt.Print("Masukkan panjang sisi miring: ")
	fmt.Scan(&sisiMiring)

	luas := 0.5 * (alas + sisiMiring) * tinggi
	fmt.Printf("Luas trapesium adalah: %.2f\n", luas)
}
