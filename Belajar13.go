package main

import "fmt"

// ==========================================
// 1. DEFINISI INTERFACE (Kontrak Perilaku)
// ==========================================

// Perbaikan: 'Type' diubah menjadi huruf kecil 'type'
// Tipe data apa pun yang memiliki method Jalan() otomatis dianggap memenuhi interface BisaBerjalan
type BisaBerjalan interface {
	Jalan()
}

// ==========================================
// 2. DEFINISI STRUCT & IMPLEMENTASI METHOD
// ==========================================

type Manusia struct {
	Nama string
}

type Robot struct {
	Kode string
}

// Struct Manusia mengimplementasikan interface BisaBerjalan karena memiliki method Jalan()
func (m Manusia) Jalan() {
	fmt.Println(m.Nama, "sedang berjalan")
}

// Struct Robot juga mengimplementasikan interface BisaBerjalan karena memiliki method Jalan()
func (r Robot) Jalan() {
	fmt.Println("Robot", r.Kode, "berjalan otomatis")
}

// ==========================================
// 3. FUNGSI YANG MENERIMA INTERFACE
// ==========================================

// Fungsi ini bersifat polimorfik: bisa menerima Manusia maupun Robot
// selama keduanya memenuhi kontrak interface BisaBerjalan
func Gerakkan(b BisaBerjalan) {
	b.Jalan()
}

// Empty Interface (interface{} atau 'any' di Go versi baru):
// Karena tidak memiliki syarat method apa pun, fungsi ini bisa menerima TIPE DATA APA SAJA
func CetakData(data interface{}) {
	fmt.Println("Data diterima:", data)
}

// ==========================================
// 4. EKSEKUSI PROGRAM
// ==========================================
func main() {
	m := Manusia{Nama: "Andi"}
	r := Robot{Kode: "RX-01"}

	// Kedua objek berbeda ini bisa dimasukkan ke fungsi yang sama
	Gerakkan(m) // Output: Andi sedang berjalan
	Gerakkan(r) // Output: Robot RX-01 berjalan otomatis

	fmt.Println("---")

	// Memanggil CetakData dengan berbagai tipe data berbeda
	CetakData("Halo Golang") // string
	CetakData(2026)          // int
	CetakData(m)             // struct Manusia
}
