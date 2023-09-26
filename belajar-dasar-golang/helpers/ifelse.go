package helpers

import "fmt"

func IfElse(){
	name := "said"

	if name == "said" {
        fmt.Println("true")
    } else if name == "alkhudri" {
		fmt.Println("else if")
	} else{
        fmt.Println("false")
    }

	// short statement
	if name2 := "joko"; len(name2) > 5{
		fmt.Println("ini terlalu panjang")
	}else{
		fmt.Println("ini tidak panjang")
	}
}