package main

import "fmt"

func main() {

	var n, faktor int

	fmt.Print("Masukkan Bilangan :")
	fmt.Scan(&n)

	faktor = 0

	if n < 2 {
		fmt.Println("Bukan Bilangan Prima")
	} else {
		for i := 1; i <= n; i++ {
			if n%i == 0 {
				faktor = faktor + 1
			} else {
				faktor = faktor
			}
		}

		if faktor == 2 {
			fmt.Println("prima")
		} else {
			fmt.Println("Bukan Bilangan Prima")
		}
	}
}

/*
Komentar (Senin, 2 Desember 2019)
*/
