package helpers

import "fmt"

// constanta tidak apa" jika tidak pernah dipakai.

// multiple constants
const (
    a = "huruf a"
    b = "huruf b"
    c = "huruf c"
)

// karena golang beranggapan mungkin akan dipanggil di tempat lain
const language = "Golang"

// functions init adalah fungsi yang akan dipanggil ketika package dipanggil
func init() {
	fmt.Println(language)
}