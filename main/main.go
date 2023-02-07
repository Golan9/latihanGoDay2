package main

import (
	"fmt"
	"latihanGoDay2/domain"
)

func main() {
	fmt.Println("-function with multiple return struct-")
	fmt.Println(domain.SetPersonalData("dam", "jakarta", 27))
	fmt.Println("")
	fmt.Println("-interface and function receiver implementation-")
	var h domain.Hitung
	h = domain.Lingkaran{10}
	fmt.Println("luas lingkaran:", h.Luas(), "keliling lingkaran:", h.Keliling())
	h = domain.Persegi{5}
	fmt.Println("luas persegi:", h.Luas(), "keliling lingkaran:", h.Keliling())
	fmt.Println("")
	domain.InterfaceKosong()
	fmt.Println("")
	fmt.Println("-function as parameter-")
	var numbers = []int{0, 1, 2, 6, 3, 7, 4, 8, 16, 22}
	var result = domain.Filter(numbers, domain.GetMoreThanSix)
	fmt.Println("angka lebih besar dari 6 :", result)
	fmt.Println("")
	domain.Test()
}
