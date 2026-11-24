package main

import "fmt"

func main() {
	// ==========================================
	// 1. ARRAY (Kapasitas tetap / fixed size)
	// ==========================================
	// Perbaikan: 'Var' diubah menjadi huruf kecil 'var'
	var angka [3]int
	angka[0] = 10
	angka[1] = 20
	angka[2] = 30
	fmt.Println("Isi array angka:", angka)

	// Perbaikan: ganti nama variabel menjadi 'angkaArray' karena 'angka' sudah dipakai
	angkaArray := [3]int{10, 20, 30}
	fmt.Println("Elemen index ke-1:", angkaArray[1]) // Output: 20

	// ==========================================
	// 2. SLICE (Kapasitas dinamis, bisa ditambah)
	// ==========================================
	// Perbaikan: ganti nama variabel menjadi 'angkaSlice'
	angkaSlice := []int{10, 20, 30}
	fmt.Println("Elemen index ke-0:", angkaSlice[0]) // Output: 10

	// Menambah elemen baru ke dalam slice menggunakan append()
	angkaSlice = append(angkaSlice, 40)
	fmt.Println("Setelah di-append:", angkaSlice) // Output: [10 20 30 40]

	// Slicing: memotong sebagian elemen dari array/slice (index awal inklusif, index akhir eksklusif)
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]                      // Ambil elemen index ke-1 (nilai 2) sampai sebelum index ke-4 (nilai 4)
	fmt.Println("Hasil slicing:", slice)   // Output: [2 3 4]

	// ==========================================
	// 3. MAP (Pasangan Key - Value)
	// ==========================================
	nilai := map[string]int{
		"Andi": 90,
		"Budi": 85,
	}

	fmt.Println("Nilai Andi:", nilai["Andi"]) // Output: 90

	// Menambah data baru & menghapus data dari map
	nilai["Citra"] = 88
	delete(nilai, "Budi")
	fmt.Println("Isi map sekarang:", nilai)

	// Comma-ok idiom: mengecek apakah sebuah key ada di dalam map atau tidak
	val, ok := nilai["Dina"]
	if ok {
		fmt.Println("Ada:", val)
	} else {
		fmt.Println("Tidak ditemukan") // Output ini tercetak karena "Dina" tidak ada di map
	}
}
