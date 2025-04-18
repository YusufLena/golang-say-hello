package golangsayhello

import "fmt"

type Age func(int) int
type Name func(string) string
func hello(name string, age int, filterName Name,  filterAge Age){
	nama := filterName(name)
	umur := filterAge(age)
	filterage := filterAge(age)
	fmt.Println("Welcome and hello Mr", nama, "Umur anda", filterage, ".", umur)
}

func NameFilter(name string) string {
	if name == "Anjing"{
		return "Nama yang anda buat sangat tidak merupakan hewan"
	} else if name == "Babi" {
		return "Nama yang anda buat sangat tidak merupakan hewan"
	} else if name == "Goblok" {
		return "Nama mengandung kata kasar"
	} else {
		return(name)
	}
}

func Umur(age int) int {
	if age < 18 {
		for i := 1; i <= 5; i++ {
			fmt.Println(i)
		}
	}
	return(age)
}

func AgeCheked(age int) string {
	if age < 18 {
		return "Umur anda Belum cukup Mr" 
	} else if age == 18 {
		return "Oke umur anda pas sekali"
	} else {
		return "umur anda sudah sangat pas"
	}
}

func SayHello(){
	filter1, filter2 := NameFilter, Umur
	hello("", 0, filter1, filter2)
}
