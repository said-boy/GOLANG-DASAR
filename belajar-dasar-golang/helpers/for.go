package helpers

import "fmt"

func ForLooping(){
	for i := 0; i < 3; i++ {
        fmt.Println(i)
    }

	slice := []string{"said", "al", "khudri"}
	
	for _ , value := range slice {
		fmt.Println(value)
	}

	datamap := map[string]string{
		"first_name": "John",
		"last_name": "Doe",
        "age": "25",
	}

	for key, value := range datamap {
		fmt.Println(key, "=", value)
	}
}