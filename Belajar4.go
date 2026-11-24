package main

import "fmt"

var versi = "1.0.0"

func main() {
	var nama string = "Golang"
	var umur int = 25
	var status bool
	var kota = "Jakarta" // otomatis bertipe string
	judul := "Belajar Golang"
	nilai := 90
	var a, b, c int = 1, 2, 3
	x, y, z := "A", "B", "C"
	var jumlah int = 10
	jumlah = 15

	fmt.Println("Versi:", versi)
	// Mencetak semua variabel lokal sekaligus agar tidak terjadi error "declared and not used"
	fmt.Println(nama, umur, status, kota, judul, nilai, a, b, c, x, y, z, jumlah)
}
