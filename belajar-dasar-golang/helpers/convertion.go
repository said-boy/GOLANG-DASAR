package helpers

import "fmt"

// hati" untuk konversi nilai yang lebih besar dari
// kapasitas dari tipe data konversinya.
var nilai32 int32 = 225
var nilai64 int64 = int64(nilai32)

var nama = "said"
var e = nama[0]
var eString string = string(e)

func Konversi()  {
	fmt.Println(nilai32)
    fmt.Println(nilai64)
    fmt.Println(eString)
}