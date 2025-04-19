package golangsayhello

func nameFilter(name string) string {
	if name == "Anjing"{
		return "Nama yang anda buat merupakan hewan"
	} else if name == "Babi" {
		return "Nama yang anda buat merupakan hewan"
	} else if name == "Goblok" {
		return "Nama mengandung kata kasar"
	} else {
		return(name)
	}
}

func ageCheked(age int) string {
	if age < 18 {
		return "Umur anda Belum cukup Mr" 
	} else if age == 18 {
		return "Oke umur anda pas sekali"
	} else {
		return "umur anda sudah sangat pas"
	}
}

func SayHello(name string, age int)(string,string, string){
	filter1 := nameFilter(name)
	filter2 := ageCheked(age)
	return filter1,"\n",filter2 + name
}
