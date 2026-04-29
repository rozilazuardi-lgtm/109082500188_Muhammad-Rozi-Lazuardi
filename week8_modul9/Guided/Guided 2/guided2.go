package main

import "fmt"

func main() {
	var nilai map[string]int = make(map[string]int)

	nilai["Rozi"] = 90
	nilai["Rafly"] = 88
	nilai["Irfan"] = 95

	fmt.Println("Data nilai : ")
	var nama string
	var grade int
	for nama, grade = range nilai {
		fmt.Println(nama, " = ", grade)
	}
	fmt.Println()

	fmt.Println("Setelah Update : ")
	nilai["Irfan"] = 92
	print("Irfan = ", nilai["Irfan"])
	fmt.Println()
	fmt.Println()

	delete(nilai, "Rozi")
	fmt.Println("Data nilai setelah delete : ")
	for nama, grade = range nilai {
		fmt.Println(nama, " = ", grade)
	}
	fmt.Println()

	fmt.Println("Hasil Searching : ")
	var cariNama string = "Rafly"
	var isiValue int
	var ok bool
	isiValue, ok = nilai[cariNama]
	if ok {
		fmt.Println("Nilai", cariNama, " = ", isiValue)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}
