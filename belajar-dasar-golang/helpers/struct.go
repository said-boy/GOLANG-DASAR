package helpers

/*
ketika membuat array dan key pada map maka tipe nya harus sama
untuk membuat data dengan berbeda tipe data gunakan struct
*/

type Customer struct {
	Name, Alamat string
	Age           int
}

func SetCustomer() Customer {
	customer := Customer{
		Name: "said",
		Alamat: "malang",
        Age: 20,
	}
	return customer
}