package helpers

import "fmt"

// dapat merubah variabel yang ada disekitarnya dengan
// nama variabel yang sama. hati"

// funciton didalam function dapat mengakses variabel
// yang ada pada parent functionnya.

// kecuali jika membuat variabel baru

func Closure() int {
	counter := 0 // (closure)
	nama := "said"
	title := "programer"
	increment := func() {
		counter++ // akan mengganti variable pada (closure)
		nama := "alkhudri"
		fmt.Println(nama) // "alkhudri"
		title = "web development"  // akan mereplace "programer"
	}

	increment() // setelah fungsi dipanggil akan merubah yang harus dirubah
	// dan menampilkan apa yang didefinisikan

	fmt.Println(title) // "web development"
	fmt.Println(nama) // "said"
	
	return counter
}