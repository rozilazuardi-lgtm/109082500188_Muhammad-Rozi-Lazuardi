package main

import "fmt"

type arrBerat [1000]float64

func main() {
	var data arrBerat
	var n int
	var i int

	fmt.Print("Masukkan jumlah kelinci: ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Print("Masukkan berat kelinci ke-", i+1, ": ")
		fmt.Scan(&data[i])
	}

	min := data[0]
	max := data[0]

	i = 1
	for i < n {
		if data[i] < min {
			min = data[i]
		}
		if data[i] > max {
			max = data[i]
		}
		i = i + 1
	}

	fmt.Println("Berat terkecil:", min)
	fmt.Println("Berat terbesar:", max)
}
