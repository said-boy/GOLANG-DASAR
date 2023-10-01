package helpers

import "strconv"

/*
parse...() -> digunakan untuk merubah string ke tipe data lain.
string -> bool, int, dll.

Format...() -> digunakan untuk merubah tipe data lain ke string
bool, int, dll -> string.
*/

func StringToBool(str string) (bool, error) {
	boolean, err := strconv.ParseBool(str) 
	return boolean, err
}

func IntToString(integer int64) string {
	number := strconv.FormatInt(integer, 10)
	return number
}