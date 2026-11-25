package main

import "fmt"

// 1. Pass by Value: mengirim salinan nilai (variabel asli di main tidak ikut berubah)
func ubahNilai(x int) {
	x = 100
}

// 2. Pass by Reference (Pointer): mengirim alamat memori (variabel asli ikut berubah)
func ubahPointer(x *int) {
	*x = 100 // Mengubah nilai yang ada di alamat memori tujuan
}

func main() {
	// Perbaikan: 'Var' diubah menjadi huruf kecil 'var'
	var angka int = 10
	// '&' (address-of) mengambil alamat memori dari variabel 'angka'
	var ptr *int = &angka

	fmt.Println("Nilai angka:", angka)
	fmt.Println("Alamat angka:", &angka)
	fmt.Println("Pointer ptr:", ptr)

	// '*' (dereference) membaca nilai asli dari alamat memori yang disimpan 'ptr'
	fmt.Println("Nilai melalui pointer:", *ptr)

	// Pengujian perbedaan Pass by Value vs Pass by Reference
	a := 10
	ubahNilai(a)
	fmt.Println("Setelah ubahNilai (tanpa pointer):", a) // Output: Tetap 10

	ubahPointer(&a)                                       // Mengirim alamat memori 'a' menggunakan '&a'
	fmt.Println("Setelah ubahPointer (pakai pointer):", a) // Output: Jadi 100

	// 3. Membuat pointer baru langsung dengan fungsi bawaan new()
	// Perbaikan: ganti nama variabel menjadi 'angkaBaru' karena 'angka' sudah dipakai di atas
	angkaBaru := new(int)
	*angkaBaru = 50
	fmt.Println("Nilai dari new(int):", *angkaBaru) // Output: 50
}
