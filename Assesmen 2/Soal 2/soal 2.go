package main

import "fmt"

const nMax int = 51

type mahasiswa struct {
	nim   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func cariNilaiPertama(T arrayMahasiswa, n int, nimCari string) int {
	for i := 1; i <= n; i++ {
		if T[i].nim == nimCari {
			return T[i].nilai
		}
	}
	return -1
}

func cariNilaiTerbesar(T arrayMahasiswa, n int, nimCari string) int {
	max := -1
	for i := 1; i <= n; i++ {
		if T[i].nim == nimCari {
			if max == -1 || T[i].nilai > max {
				max = T[i].nilai
			}
		}
	}
	return max
}

func main() {
	var T arrayMahasiswa
	var n int
	var nimCari string

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&n)

	if n > nMax-1 {
		n = nMax - 1
	}

	for i := 1; i <= n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i)
		fmt.Scan(&T[i].nim, &T[i].nama, &T[i].nilai)
	}

	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&nimCari)

	pertama := cariNilaiPertama(T, n, nimCari)
	terbesar := cariNilaiTerbesar(T, n, nimCari)

	if pertama != -1 {
		fmt.Printf("Nilai pertama dari nim %s adalah %d\n", nimCari, pertama)
		fmt.Printf("Nilai terbesar dari nim %s adalah %d\n", nimCari, terbesar)
	} else {
		fmt.Println("Data mahasiswa dengan nim tersebut tidak ditemukan.")
	}
}
