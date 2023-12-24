package helpers

import "strings"

// package untuk memanipilasi string

func DeleteSpaceInFrontAndBack(message *string) string {
	msg := strings.Trim(*message, " ") 
	return msg
}