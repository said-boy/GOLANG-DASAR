package helpers

import "fmt"

func ArraysTo(){
	// cara manual
	var names [3]string
	names[0] = "John"
	names[1] = "Mary"
	names[2] = "Peter"

	// secara singkat
	var nilai = [3]int{80,90,100}
	
	nilai[2] = 75 // merubah data pada array 100 -> 75

	fmt.Println(names)
	fmt.Println(nilai)
	fmt.Println(len(nilai)) // 3 banyak daya tampung pada array

	var nilai2 = [10]int{80,90,100}
	fmt.Println(len(nilai2)) // 10 banyak daya tampung pada array

	/*
	gunakan [...] memberikan kebebasan untuk menambahkan
	data array sebanyak yang anda mau. tidak sepesifik
	*/ 
	var nilai3 = [...]int{100,200,300,400,500}
	fmt.Println(nilai3) 
}