package helpers

import "errors"

func Pembagian(a int, b int) (int, error){
	if b == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		return a / b, nil
	}
}


