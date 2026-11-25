package main

import "fmt"

// 1. Fungsi tanpa parameter dan tanpa return value
// Perbaikan: 'Func' diubah menjadi huruf kecil 'func'
func sapa() {
	fmt.Println("Halo dari fungsi!")
}

// 2. Fungsi dengan satu parameter
func halo(nama string) {
	fmt.Println("Halo", nama)
}

// 3. Fungsi dengan beberapa parameter bertipe sama
// Penulisan singkat (a, b int) sama artinya dengan (a int, b int).
// Salah satu fungsi 'tambah' dihapus agar tidak terjadi error duplikat nama fungsi.
func tambah(a, b int) {
	fmt.Println("Hasil tambah:", a+b)
}

// 4. Fungsi dengan satu return value (mengembalikan nilai)
func kali(a, b int) int {
	return a * b
}

// 5. Fungsi dengan multiple return values (mengembalikan lebih dari satu nilai)
func hitung(a, b int) (int, int) {
	return a + b, a * b
}

// 6. Fungsi dengan Named Return Value (variabel return sudah dinamai di awal)
func bagi(a, b int) (hasil int) {
	hasil = a / b
	return // 'naked return': otomatis mengembalikan nilai variabel 'hasil'
}

// Semua pemanggilan fungsi wajib dilakukan di dalam func main()
func main() {
	sapa()
	halo("Golang")
	tambah(10, 5)

	hasil := kali(3, 4)
	fmt.Println("Hasil kali:", hasil) // Output: 12

	jumlah, perkalian := hitung(2, 5)
	fmt.Println("Jumlah:", jumlah)       // Output: 7
	fmt.Println("Perkalian:", perkalian) // Output: 10

	hasilBagi := bagi(20, 4)
	fmt.Println("Hasil bagi:", hasilBagi) // Output: 5

	// 7. Anonymous Function (Fungsi tanpa nama yang disimpan ke dalam variabel)
	f := func(nama string) {
		fmt.Println("Hai", nama)
	}
	f("Golang")
}
