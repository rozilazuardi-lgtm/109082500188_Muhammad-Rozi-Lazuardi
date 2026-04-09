package main

import "fmt"

const pi float64 = 3.14

func volume(r, t float64) float64 {
	return pi * r * r * t
}

func massa(r, t, p float64) float64 {
	v := volume(r, t)
	return v * p
}

func display(m1, m2 float64) {
	if m1 == m2 {
		fmt.Println("BALANCE")
	} else {
		selisih := m1 - m2
		if selisih < 0 {
			selisih = -selisih
		}
		fmt.Println("selisih massa zat cair kiri dan massa jenis zat cair kanan:", selisih)
	}
}

func main() {
	var r float64
	var tKiri, tKanan float64
	var mjKiri, mjKanan float64
	var massaKiri, massaKanan float64

	fmt.Print("masukkan jari-jari alas tabung: ")
	fmt.Scan(&r)

	fmt.Print("masukkan tinggi zat cair di tabung kiri: ")
	fmt.Scan(&tKiri)
	fmt.Print("masukkan massa jenis zat cair di tabung kiri: ")
	fmt.Scan(&mjKiri)

	fmt.Print("masukkan tinggi zat cair di tabung kanan: ")
	fmt.Scan(&tKanan)
	fmt.Print("masukkan massa jenis zat cair di tabung kanan: ")
	fmt.Scan(&mjKanan)

	massaKiri = massa(r, tKiri, mjKiri)
	massaKanan = massa(r, tKanan, mjKanan)

	display(massaKiri, massaKanan)
}
