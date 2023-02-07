package domain

import (
	"fmt"
	"math"
)

func (l Lingkaran) Luas() float64 {
	jariJari := l.Diameter / 2
	return math.Pi * math.Pow(jariJari, 2)
}

func (l Lingkaran) Keliling() float64 {
	return math.Pi * l.Diameter
}

func (p Persegi) Luas() float64 {
	return p.Sisi * p.Sisi
}

func (p Persegi) Keliling() float64 {
	return 4 * p.Sisi
}

// interface kosong
func InterfaceKosong() {
	fmt.Println("-interface kosong-")
	var k Kosong
	k = "dam"
	fmt.Println("interface string:", k)
	k = 12.4
	var i float64 = 5.5
	result := i * k.(float64)
	fmt.Println("interface int:", k)
	fmt.Println("interface float:", result)
	k = []string{"dam", "dim"}
	fmt.Println("interface slice:", k)
}

// -function as parameter-
func GetMoreThanSix(number int) bool {
	return number > 6
}

func Filter(numbers []int, callback func(int) bool) (result []int) {
	for _, value := range numbers {
		if filtered := callback(value); filtered {
			result = append(result, value)
		}
	}
	return result
}

// function with multiple return
func SetPersonalData(n string, l string, a int) (PersonalData, error) {
	var p PersonalData
	p.Name = n
	p.Location = l
	p.Age = a
	if p.Name == "" || p.Location == "" || p.Age == 0 {
		return p, fmt.Errorf("salah satu isian tidak boleh kosong")
	}
	return p, nil
}

func Test() {
	fmt.Println("hello world")
}
