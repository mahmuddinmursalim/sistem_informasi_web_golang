package main

import "fmt"

func main() {
	// 1. Percabangan If - Else If - Else
	// Perbaikan: ubah 'Nilai' menjadi 'nilai' (huruf kecil) agar sesuai dengan kondisi di bawah
	nilai := 85

	if nilai >= 90 {
		fmt.Println("Nilai A")
	} else if nilai >= 80 {
		fmt.Println("Nilai B") // Output ini yang akan tercetak karena 85 >= 80
	} else {
		fmt.Println("Nilai C atau di bawahnya")
	}

	// 2. Percabangan Switch Case
	hari := "Senin"

	switch hari {
	case "Senin":
		fmt.Println("Awal minggu") // Output ini yang tercetak
	case "Sabtu", "Minggu": // Bisa mencocokkan lebih dari satu nilai sekaligus
		fmt.Println("Weekend!")
	default:
		fmt.Println("Hari biasa")
	}

	// 3. Perulangan (For Loop) standar
	// Mencetak angka dari 0 sampai 4
	for i := 0; i < 5; i++ {
		fmt.Println("Loop ke:", i)
	}

	// 4. Infinite Loop (Perulangan tanpa henti) dengan 'break'
	for {
		fmt.Println("Terus jalan...")
		break // Wajib ada 'break' atau 'return' agar program tidak terjebak selamanya
	}

	// 5. Perulangan For-Range untuk membaca isi Slice/Array
	angka := []int{1, 2, 3, 4}

	// Mengambil index (posisi) dan value (nilai) sekaligus
	for index, value := range angka {
		fmt.Println("Index:", index, "Value:", value)
	}

	// Menggunakan '_' (blank identifier) untuk mengabaikan index jika hanya butuh nilainya
	for _, value := range angka {
		fmt.Println("Hanya value:", value)
	}
}
