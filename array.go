package main

import "fmt"

func main() {

	var countries [5]string
	odd_numbers := [5]int{1, 3, 5, 7, 9}
	test := [...]int{2, 3, 3}
	even_numbers := [5]int{1: 2, 3: 5}
	number := [5]string{"a", "b", "c"}

	countries[0] = "india"
	countries[1] = "test"

	fmt.Println(countries[0])
	fmt.Println(countries[1])

	fmt.Println(odd_numbers)

	fmt.Println(test)

	fmt.Println(even_numbers)

	for index, value := range number {
		fmt.Println(index, " => ", value)
	}
}

/*
Komentar (Selasa, 3 Desember 2019)
*/
