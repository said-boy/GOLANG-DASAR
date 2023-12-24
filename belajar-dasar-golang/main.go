package main

import (
	"belajar-dasar-golang/helpers"
	"belajar-dasar-golang/myinterface"
	"fmt"
)

func startFunction() {
	fmt.Println("Belajar Golang function internal")
}

func main() {
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

	arr := [3]string{"a", "b", "c"}
	helpers.ArrayParameter(arr)

	fmt.Println(helpers.ReturnValue())
	first_name, last_name := helpers.ReturnMultipleValues()
	fmt.Println(first_name, last_name)
	fmt.Println(helpers.ReturnNameValue())

	fmt.Println(helpers.Variadic(10, 10, 10))
	data := []int{10, 10, 10}
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
		Name:   "Lancelot",
		Role:   "Assassin",
		Blood:  100,
		Attack: 200,
	}

	ruby := helpers.Person{
		Name:   "Ruby",
		Role:   "Tank",
		Blood:  100,
		Attack: 200,
	}
	helpers.PrintStrengthAndPunch(lancelot)
	helpers.PrintStrengthAndPunch(ruby)

	// interface kosong
	fmt.Println(helpers.Message("halo", "apa", "kabar", 1, true))

	// nil
	fmt.Println(helpers.PersonNil) // slice nil

	// interface error
	hasil, err := helpers.Pembagian(10, 0)
	if err != nil { // tidak ada error
		fmt.Println(err)
	} else {
		fmt.Println(hasil)
	}

	// type assertion
	/*
		digunakan untuk interface kosong
		harus yakin jika data yang dikembalikan berpa data yang sama
		jika berbeda maka akan panic
	*/
	random := helpers.Random()
	fmt.Println(random.(string)) // harus sama

	switch value := random.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("Unknown", value)
	}

	// pointer
	orang1 := helpers.Address{
		Kota:     "Pekalongan",
		Provinsi: "Jawa Tengah",
	}
	orang2 := &orang1
	orang2.Kota = "Kendal" // orang1 juga berubah

	// orang3 := &orang1
	// *orang3 = helpers.Address{
	// 	Kota: "Jogja",
	// 	Provinsi: "Jawa timur",
	// }

	fmt.Println(orang1)
	fmt.Println(orang2)
	// fmt.Println(orang3)

	// pointer in function
	// contoh data aslinya.
	user1 := helpers.Akun{
		Username: "user1",
		Password: "WeekPassword",
	}

	helpers.ChangePasswordNoPointer(user1) // tidak meruubah data aslinya
	fmt.Println(user1)                     // WeekPassword

	helpers.ChangePassword(&user1) // merubah data aslinya
	fmt.Println(user1)             // SecurePassword

	user2 := &user1
	user2.Username = "user2"

	fmt.Println(user1)
	fmt.Println(*user2)

	// package os
	hostname, err := helpers.PrintHostname()
	if err == nil {
		fmt.Println("Hostname: ", hostname)
	} else {
		fmt.Println(err.Error())
	}

	// flag
	Host, User, Password := helpers.DataConnectionDatabase()

	fmt.Println("Host : ", *Host)
	fmt.Println("User : ", *User)
	fmt.Println("Password : ", *Password)

	// strings
	message := "    Please Subscribe    "
	message2 := &message
	message = helpers.DeleteSpaceInFrontAndBack(&message)
	fmt.Println(message)
	fmt.Println(*message2)

	// parse string to Bool strconv
	boolean, err := helpers.StringToBool("0")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number := helpers.IntToString(100)
	fmt.Println(number)

	fmt.Println(helpers.Round)
	fmt.Println(helpers.Floor)

	// memanggil function dari interfacenya. (sesuai dengan golang)
	var Hero myinterface.IHero = &myinterface.Hero{}
	fmt.Println(Hero.Run(30))

	var Enemy myinterface.IHero = &myinterface.Enemy{}
	fmt.Println(Enemy.Run(30))

	// memanggil function dari structnya. (meningkatkan kecepatan)
	var Hero1 myinterface.Hero = myinterface.Hero{Name: "Said"}
	fmt.Println(Hero1.PrintName())
	fmt.Println(Hero1.Run(20))

}
