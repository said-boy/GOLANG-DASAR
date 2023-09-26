package helpers

/*

nama function diawali huruf kecil tidak bisa diakses dari tempat lain
nama function diawali huruf besar dapat diakses dari tempat lain 

*/

func SayHello() string {
	return "Hello Said!"	
}

// tidak bisa diakses dari main.go
func sayHello2() string {
	return "Hello Said!"	
}