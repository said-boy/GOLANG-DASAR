package main

import (
	"belajar-dasar-golang/helpers"
	"fmt"
)

func startFunction(){
	fmt.Println("Belajar Golang function internal")
}

func main(){
	fmt.Println(helpers.SayHello())
	startFunction()
	helpers.Nama()
	fmt.Println(helpers.Boolean())
	helpers.PanjangNama()
	helpers.Konversi()
	helpers.Alias()
	helpers.Matematika()
	helpers.Perbandingan()
	helpers.ArraysTo()
	fmt.Println("ini slice")
	helpers.Slice()
	helpers.Map()
	helpers.IfElse()
	helpers.Switch()
	helpers.ForLooping()
	helpers.Break()
	helpers.FuncParameter("Boy", 20)

	arr := [3]string{"a","b","c",}
	helpers.ArrayParameter(arr)

	fmt.Println(helpers.ReturnValue())
	first_name, last_name := helpers.ReturnMultipleValues()
	fmt.Println(first_name, last_name)
	fmt.Println(helpers.ReturnNameValue())

	fmt.Println(helpers.Variadic(10,10,10))
	data := []int{10,10,10}
	// jika datanya sudah ada tambahkkan ... dibelakangnya
	fmt.Println(helpers.VariadicWithData(data...))

	// menjadi alias
	sayGoodbye := helpers.GetGoodBye
	fmt.Println(sayGoodbye("Kemalasan"))

	status := helpers.SayStatus(20, helpers.GetStatus)
	fmt.Println(status)

	// function anonymous
	user := "Admin"
	isAdmin := helpers.IsAdmin(user, func(string) bool {
		return user == "Admin"
	})
	fmt.Println(isAdmin)


}