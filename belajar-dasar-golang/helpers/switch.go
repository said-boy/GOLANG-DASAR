package helpers

import "fmt"

func Switch() {

	name := "sample"

	switch name {
	case "sample":
		fmt.Println("sample")
    default:
		fmt.Println("not sample")
	}

	//short statement
	switch nama2 := "said"; len(nama2) < 5{
	case true:
		fmt.Println("nama terlalu pendek")
	case false:
		fmt.Println("nama pass")
	}

	length := len(name)
	switch {
	case length < 5:
		fmt.Println("nama terlalu pendek")
	case length > 5 && length < 10:
		fmt.Println("nama pass")
	default:
		fmt.Println("nama terlalu panjang")
	}

}