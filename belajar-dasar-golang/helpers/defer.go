// defer adalah fungsi yang akan
// dipanggil ketika sebuah fungsi
// selesai dijalankan.
// walaupun fungsi tersebut menghasilkan error

package helpers

import "fmt"

func logging(){
	fmt.Println("Program berakhir")
}

func RunApplication(){
	defer logging() // biasakan menggunakan defer itu diatas
	fmt.Println("Program mulai")
}