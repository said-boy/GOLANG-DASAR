package helpers

// pointer , length, capacity

import (
    "fmt"
)

func Slice() {
    /*
    array bisa diubah datanya. tetapi tidak bisa
    ditambah. jika saya punya array dengan
    panjang 5 maka saya tidak bisa menambahkan
    data array ke 6. Error!.
    */
    var array = [5]int{1, 2, 3, 4, 5}

    /*
    untuk dapat memanipulasi array gunakan slice
    (slice akan sering digunakan dibanding array)
    */ 

    /* 
    mengambil potongan data dari array
    secara realtime 
    */
    var slice = array[2:4] 
    fmt.Println(slice) // [4,5]

    /*
    jadi ketika data pada slice diubah maka sata pada
    array juga akan ikut berubah dengan catatan
    data yang dirubah masih dalam cakupan data
    pada array.
    */
    slice[0] = 7
    fmt.Println(array) // [1,2, 7, 4, 5]
    
    /*
    append()
    akan mengganti data paling belakang jika masih 
    cukup. 
    */
    slice = append(slice, 8)
    fmt.Println(slice) // [7, 4, 8]
    fmt.Println(array) // [1,2, 7, 4, 8]

    /*
    tetapi akan membuat variable baru jika
    tempat pada array tidak cukup
    */
    slice = append(slice, 9, 10)
    
    /*
    karena membuat variable baru maka array tidak
    akan terpengarug oleh perubahan pada slice
    */
    fmt.Println(slice) // [7, 4, 8, 9, 10] ini variabel baru
    fmt.Println(array) // [1,2, 7, 4, 8] tidak berubah
    
    /*
    Hati" berbeda sedikit dapat mengecoh.
    array harus mendefinisikan banyak data nya
    walaupun dengan [...] atau sepesefik [3]

    slice tidak, cukup gunakana [] saja.
    */
    iniArray := [...]string{"said", "al", "khudri"}
    fmt.Println(iniArray)

    iniSlice := []string{"said", "al", "khudri"}
    fmt.Println(iniSlice)

    /*
    gunakan make() untuk membuat slice dengan
    lebih aman.
    parameternya: make([]Type data, length, capacity)
    
    length = banyak data yang akan diisikan
    capacity = banyak data yang bisa ditampung
    */
    makeSlice := make([]string, 2, 7)
    makeSlice[0] = "Hello"
    makeSlice[1] = "World!"
    fmt.Println(makeSlice)
    
    // Memanbah data ke akhir
    makeSlice = append(makeSlice, "Baru")
    fmt.Println(makeSlice)

}