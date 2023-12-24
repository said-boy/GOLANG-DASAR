package helpers

func CustomerSay() Customer {
	customer := Customer{
		Name: "said",
		Alamat: "malang",
        Age: 20,
	}
	return customer
}

func (client Customer) SayHi(name string) string{
	return "Hi " + name + ", Nama saya : " + client.Name
}

