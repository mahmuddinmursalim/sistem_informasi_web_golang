package main

import "fmt"

// ==========================================
// 1. DEFINISI STRUCT (Cetakan Data)
// ==========================================

// Struct dasar
// Perbaikan: 'Type' diubah menjadi huruf kecil 'type'
type Mahasiswa struct {
	Nama  string
	Umur  int
	Aktif bool
}

// Struct tambahan untuk contoh Nested Struct (Struct di dalam Struct)
type Alamat struct {
	Kota  string
	Jalan string
}

// Perbaikan: Nama struct dibedakan menjadi 'MahasiswaLengkap' karena nama 'Mahasiswa' sudah dipakai di atas
type MahasiswaLengkap struct {
	Nama   string
	Umur   int
	Alamat Alamat // Menggabungkan struct Alamat sebagai tipe data field
}

// ==========================================
// 2. FUNGSI DENGAN POINTER STRUCT
// ==========================================

// Menerima pointer (*Mahasiswa) agar perubahan field Umur berdampak langsung ke data aslinya
func ubahUmur(m *Mahasiswa) {
	// Di Go, kita tidak perlu menulis (*m).Umur, cukup m.Umur saja (otomatis di-dereference)
	m.Umur = 25
}

// ==========================================
// 3. EKSEKUSI PROGRAM DI DALAM MAIN
// ==========================================
func main() {
	// Cara 1: Deklarasi variabel dulu, lalu isi field satu per satu
	var mhs1 Mahasiswa
	mhs1.Nama = "Budi"
	mhs1.Umur = 20
	mhs1.Aktif = true
	fmt.Println("Data mhs1:", mhs1)

	// Cara 2: Deklarasi sekaligus inisialisasi langsung (Struct Literal)
	mhs2 := Mahasiswa{
		Nama:  "Ani",
		Umur:  22,
		Aktif: false,
	}
	fmt.Println("Nama mhs2:", mhs2.Nama)
	fmt.Println("Umur mhs2:", mhs2.Umur)

	// Cara 3: Mengisi Nested Struct
	mhs := MahasiswaLengkap{
		Nama: "Citra",
		Umur: 21,
		Alamat: Alamat{
			Kota:  "Bandung",
			Jalan: "Jl. Merdeka",
		},
	}
	// Mengakses field di dalam struct bersarang menggunakan tanda titik berantai
	fmt.Println("Kota Citra:", mhs.Alamat.Kota) // Output: Bandung

	// Cara 4: Mengubah isi struct menggunakan fungsi pointer
	m := Mahasiswa{Nama: "Doni", Umur: 20}
	ubahUmur(&m)                                 // Kirim alamat memori 'm' menggunakan tanda '&'
	fmt.Println("Umur Doni setelah diubah:", m.Umur) // Output: 25
}
