package golangsayhello

import "fmt"

type Age func(int) int
type Name func(string) string
func Hello(name string, filterName Name, age int, filterAge Age){
	nama := filterName(name)
	umur := filterAge(age)
	fmt.Println("Welcome and hello Mr", nama, "Umur anda", filterAge, ".", umur)
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

func AgeCheked(umur int) string {
	if umur < 18 {
		return "Umur anda Belum cukup Mr" 
	} else if umur == 18 {
		return "Oke umur anda pas sekali"
	} else {
		return "umur anda sudah sangat pas"
	}
}

