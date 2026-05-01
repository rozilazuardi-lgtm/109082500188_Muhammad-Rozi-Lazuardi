package main

import "fmt"

type arrBerat [1000]float64

func main() {
	var data arrBerat
	var x, y int
	var i int

	fmt.Print("Masukkan jumlah ikan (x): ")
	fmt.Scan(&x)

	fmt.Print("Masukkan kapasitas per wadah (y): ")
	fmt.Scan(&y)

	for i = 0; i < x; i++ {
		fmt.Print("Masukkan berat ikan ke-", i+1, ": ")
		fmt.Scan(&data[i])
	}

	var jumlahWadah int
	if x%y == 0 {
		jumlahWadah = x / y
	} else {
		jumlahWadah = (x / y) + 1
	}

	var totalPerWadah float64
	var totalSemua float64
	var count int = 0

	fmt.Println("Total berat tiap wadah:")

	for i = 0; i < x; i++ {
		totalPerWadah = totalPerWadah + data[i]
		count = count + 1

		if count == y || i == x-1 {
			fmt.Println(totalPerWadah)
			totalSemua = totalSemua + totalPerWadah

			totalPerWadah = 0
			count = 0
		}
	}

	rata := totalSemua / float64(jumlahWadah)

	fmt.Println("Rata-rata berat per wadah:", rata)
}
