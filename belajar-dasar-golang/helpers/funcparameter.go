package helpers

import "fmt"

func FuncParameter(nama string, umur int){
	fmt.Println("nama",nama,"umur",umur)
}

func ArrayParameter(arr [3]string){
	for _, value := range arr {
		fmt.Println(value)
	}
}