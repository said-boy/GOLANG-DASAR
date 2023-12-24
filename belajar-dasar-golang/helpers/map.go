package helpers

import "fmt"

func Map(){
	// data key value

	person := map[string]string{
		"name": "said",
        "age": "20",
	}
	// menambahkan data
	person["title"] = "Programer"
	
	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person["title"])
	
	// mengubah data
	person["age"] = "21"
	fmt.Println(person["age"])
	
	fmt.Println(len(person))

} 