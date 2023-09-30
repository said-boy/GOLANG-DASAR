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

	fmt.Println(helpers.Fibonancy(5))
	fmt.Println(helpers.Closure())
	helpers.RunApplication()

	// panic
	helpers.PanicApplication(false)
	
	// recovery
	helpers.RunApp(true)
	fmt.Println("Tes dapat diesekusi")

	// struct
	fmt.Println(helpers.SetCustomer())
	//struct function
	people := helpers.SetCustomer()
	fmt.Println(people.SayHi("Pelanggan"))

	// Geng (Interface)
	lancelot := helpers.Person{
		Name: "Lancelot",
		Role: "Assassin",
		Blood: 100,
		Attack: 200,
	}
	
	ruby := helpers.Person{
		Name: "Ruby",
		Role: "Tank",
		Blood: 100,
		Attack: 200,
	}
	helpers.PrintStrengthAndPunch(lancelot)
	helpers.PrintStrengthAndPunch(ruby)

	// interface kosong
	fmt.Println(helpers.Message("halo","apa","kabar",1,true))

	// nil
	fmt.Println(helpers.PersonNil) // slice nil

	// interface error
	hasil, err := helpers.Pembagian(10,0)
	if err != nil { // tidak ada error
        fmt.Println(err)
	} else {
        fmt.Println(hasil)
	}


}