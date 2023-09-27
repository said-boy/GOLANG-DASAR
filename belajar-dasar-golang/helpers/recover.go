package helpers

import "fmt"

// jika ada panic maka program akan berhenti dan akan mengembalikan error
// recover adalah fungi yang menangani atau menampung pesan error tersebut
// saat panic, jadi program tidak mengembalikan terlalu panjang error
// dan lebih rapih saat ada serangan panic.
// dan recover ini akan membuat program tetap berjalan walaupun error terjadi


func endProgramWithRecovery(){
	fmt.Println("Program Selesai...")
	message := recover() 
	// reccover harus diletakkan pada defer function, karena ketika panic
	// program tidak akan melanjutkan eksekusi dan akan kembali keatas.
	// dan ketika panic bergerak keatas akan bertemu dengan defer function.
	// dimana didalam defer function terdapat recover yang akan mengangkap pesan error dari panic yang terjadi.
	if message != nil{
        fmt.Println("Error:", message)
    }
}

func RunApp(error bool){
	defer endProgramWithRecovery() // defer function
	if error {
		panic("Program Mengalami Error") // pesan ini akan ditangkap oleh recover() yang ada di defer function
	}
	for i := 1; i < 5; i++ {
		fmt.Println("Program Berjalan...")
	}
}