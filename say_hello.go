package golangsayhello

import (
	"fmt"

	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

func endApp(){
	fmt.Println("menutup apk")
}

func SayHello(name string) string {
	defer endApp()
	return "Hello" + name
}

func Iterasi(name string) string {
	defer endApp()
	if name != "" {
		for i := 0; i < len(name); i++ {
			return "Welcome" + name
		}
	} else{
		fmt.Println("Name not found")
	}
	return name
}