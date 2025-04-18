package golangsayhello

import "fmt"

func endApp(){
	fmt.Println("menutup apk")
}

func SayHello(name string) string {
	defer endApp()
	return "Hello" + name
}

func Iterasi(name string) {
	defer endApp()
	if name != "" {
		for i := 0; i < len(name); i++ {
			fmt.Println("Welcome", name)
			break
		}
	} else{
		fmt.Println("Name not found")
	}
}