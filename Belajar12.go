package main

import "fmt"

/*
	POLA DASAR PENULISAN METHOD DI GOLANG:
	func (r ReceiverType) namaMethod() {
		// isi method
	}
	*Catatan: kata kunci 'func' wajib huruf kecil semua.
*/

// 1. Definisi Struct sebagai pemilik method
type Mahasiswa struct {
	Nama string
	Umur int
}

// 2. Method dengan Value Receiver (m Mahasiswa)
// Hanya membaca data struct tanpa mengubah nilai aslinya
func (m Mahasiswa) Perkenalan() {
	fmt.Printf("Halo, nama saya %s dan umur saya %d\n", m.Nama, m.Umur)
}

// 3. Method dengan Parameter dan Return Value
func (m Mahasiswa) TahunLahir(tahunSekarang int) int {
	return tahunSekarang - m.Umur
}

// 4. Value Receiver vs Pointer Receiver saat mengubah data
// VALUE RECEIVER (m Mahasiswa): hanya mengubah salinan (copy) data, data asli TIDAK berubah
func (m Mahasiswa) UbahNamaBaru(nama string) {
	m.Nama = nama
}

// POINTER RECEIVER (m *Mahasiswa): menggunakan tanda '*', mengubah data asli di memori
func (m *Mahasiswa) UbahNama(nama string) {
	m.Nama = nama
}

// Perbaikan: Menggabungkan kedua blok 'main()' menjadi satu fungsi main saja
func main() {
	// Pengujian 1: Memanggil method Perkenalan dan TahunLahir
	mhs1 := Mahasiswa{Nama: "Adit", Umur: 21}
	mhs1.Perkenalan()
	fmt.Println("Tahun lahir Adit:", mhs1.TahunLahir(2026))

	fmt.Println("---")

	// Pengujian 2: Perbandingan Value Receiver vs Pointer Receiver
	mhs2 := Mahasiswa{Nama: "Rudi", Umur: 22}

	// Memanggil UbahNamaBaru (Value Receiver) -> Nama tetap "Rudi"
	mhs2.UbahNamaBaru("Budi")
	fmt.Println("Setelah UbahNamaBaru (tanpa pointer):", mhs2.Nama) // Output: Rudi

	// Memanggil UbahNama (Pointer Receiver) -> Nama berubah menjadi "Santi"
	mhs2.UbahNama("Santi")
	fmt.Println("Setelah UbahNama (pakai pointer):", mhs2.Nama) // Output: Santi
}
