package main

import "fmt"

func main() {
	var usia int = 25
	var tinggi float32 = 175.5
	var nama string = "Golang"
	pesan := `Halo,
Ini adalah teks
dengan banyak baris.`
	var aktif bool = true

	// Perbaikan 1: Gunakan '=' (bukan ':=') karena variabel 'nama' sudah ada
	nama = "Belajar Golang"

	umur := 20
	lulus := true
	var x int = 10
	var y float64 = float64(x)

	var angka int
	fmt.Println(angka) // Output: 0

	// Perbaikan 2: Gunakan semua variabel yang sudah dibuat agar tidak error
	fmt.Println("Usia:", usia, "| Umur:", umur)
	fmt.Println("Tinggi:", tinggi)
	fmt.Println("Nama:", nama)
	fmt.Println("Pesan:", pesan)
	fmt.Println("Aktif:", aktif, "| Lulus:", lulus)
	fmt.Println("Nilai x & y:", x, y)

	fmt.Println("Hello, World!")
}

//Yang bawah ini adalah print versi semua
package main

import "fmt"

func main() {
	var usia int = 25
	var tinggi float32 = 175.5
	var nama string = "Golang"
	pesan := `Halo,
Ini adalah teks
dengan banyak baris.`
	var aktif bool = true

	nama = "Belajar Golang"
	umur := 20
	lulus := true
	var x int = 10
	var y float64 = float64(x)
	var angka int

	// Langsung print semua variabel sekaligus dalam satu baris kode
	fmt.Println(usia, tinggi, nama, pesan, aktif, umur, lulus, x, y, angka, "Hello, World!")
}
