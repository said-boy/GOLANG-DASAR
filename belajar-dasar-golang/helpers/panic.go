package helpers

import "fmt"

// panic function adalah untuk menghentikan fungsi
// apabila terjadi error/kesalahan pada fungsi


func endApplication(){
	fmt.Println("Aplikasi selesai")
}

func PanicApplication(error bool){
	defer endApplication()
	if error{
		panic("Aplikasi gagal")
	}
	fmt.Println("Aplikasi berhasil")
}
