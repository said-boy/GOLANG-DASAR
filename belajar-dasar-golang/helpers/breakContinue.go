package helpers

import "fmt"

func Break(){
	for i := 1; i < 3; i++ {
		fmt.Println(i)
		if i == 2 {
			break
		}
    }

	for i := 1; i < 4; i++ {
		fmt.Println(i)
		if i == 2 {
			continue
		}
    }

	/* 
	outerloop akan menghnetikan semua perulangan
	*/
	paksaberhenti:
	for i := 1; i < 4; i++ {
        for j := 1; j < 4; j++ {
            fmt.Println(i, j)
			if j == 2 {
				break paksaberhenti
			}
        }
    }
	
}