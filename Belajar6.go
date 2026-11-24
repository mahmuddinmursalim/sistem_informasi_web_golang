package main

import "fmt"

func main() {
	// 1. Operator Aritmatika
	a := 10 // Perbaikan: gunakan 'a' huruf kecil agar sesuai dengan pemanggilan di bawah
	b := 3
	fmt.Println("Penjumlahan (a + b):", a+b) // Output: 13
	fmt.Println("Modulus / Sisa Bagi (a % b):", a%b) // Output: 1

	// 2. Operator Perbandingan (menghasilkan nilai boolean: true/false)
	x := 5
	y := 7
	fmt.Println("x < y:", x < y)   // Output: true (5 lebih kecil dari 7)
	fmt.Println("x == y:", x == y) // Output: false (5 tidak sama dengan 7)

	// 3. Operator Logika
	// Perbaikan: ganti nama variabel menjadi 'p' dan 'q' karena 'a' dan 'b' sudah terpakai di atas
	p := true
	q := false
	fmt.Println("p && q (AND):", p && q) // Output: false (keduanya harus true agar hasilnya true)
	fmt.Println("p || q (OR):", p || q)  // Output: true (salah satu true saja hasilnya true)
	fmt.Println("!p (NOT):", !p)         // Output: false (membalikkan nilai true menjadi false)
}
