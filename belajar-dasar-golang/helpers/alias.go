package helpers

import "fmt"

// membuat tipe data sendiri
// dari tipe yg sudah ada data golang

func Alias(){
	type NoKtp string
	type Umur int

	var ktp NoKtp = "208383969" 
	var umur Umur = 20 
	fmt.Println(ktp)
	fmt.Println(umur)

}



