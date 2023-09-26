package helpers

import "fmt"

// Error contoh (3)
// contohSalah := "salah" 

func Nama(){
	// macam" penulisan variable

	// 1. menggunakan var dan menyebutkan tipe datanya
	// var nama string
	// nama = "said"
	
	// 2. menggunakan var tanpa menyebutkan tipe datanya
	// var nama = "said"
	
	// 3. tanpa var dan tanpa menyebutkan tipe datanya
	// diawal pertama kali membuat variabel ( pengganti var )
	// dan hanya dapat didalam function
	nama := "said"
	nama = "alkhudri" // perubahan selanjutnya cukup = saja.

	fmt.Println(nama)

	// mutiple variable
	var (
		umur = 20
		gender = "male"
	)
	fmt.Println(umur, gender)

}
