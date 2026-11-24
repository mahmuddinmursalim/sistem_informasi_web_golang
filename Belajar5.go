package main

import "fmt"

/*
	Konstanta (const) adalah tempat menyimpan nilai yang bersifat tetap
	dan TIDAK BISA diubah lagi selama program berjalan.
*/

// 1. Konstanta tunggal tanpa penulisan tipe data eksplisit (Untyped Constant)
const phi = 3.14
const aplikasi = "Belajar Golang"

// 2. Perbedaan Untyped dan Typed Constant
const angka = 10    // Untyped: tipe datanya fleksibel, menyesuaikan konteks saat dipakai
const umur int = 25 // Typed: tipe datanya dikunci mati sebagai bilangan bulat (int)

// 3. Deklarasi banyak konstanta sekaligus dalam satu blok tanda kurung
const (
	satu = 1
	dua  = 2
	tiga = 3
)

// 4. Penggunaan 'iota' untuk membuat urutan angka otomatis (dimulai dari 0)
const (
	A = iota // Bernilai 0
	B        // Otomatis bertambah 1 -> bernilai 1
	C        // Otomatis bertambah 1 -> bernilai 2
)

// func main() wajib ada sebagai pintu masuk utama saat program dijalankan
func main() {
	// Mencetak semua konstanta ke layar menggunakan package fmt
	fmt.Println("Aplikasi:", aplikasi, "| Phi:", phi)
	fmt.Println("Angka:", angka, "| Umur:", umur)
	fmt.Println("Kelompok angka:", satu, dua, tiga)
	fmt.Println("Hasil iota:", A, B, C)

	// Contoh bukti konstanta tidak bisa diubah:
	// phi = 3.15 // Jika baris ini diaktifkan, Go akan error: "cannot assign to phi"
}
