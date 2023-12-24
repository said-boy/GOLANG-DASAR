package helpers

import (
	"flag"
)

/*
untuk membuat -host -user -password
*/ 

func DataConnectionDatabase() (*string, *string, *string) {
	Host := flag.String("host", "localhost", "Put your host")
	User := flag.String("user", "root", "Put your username")
	Password := flag.String("password", "root", "Put your password")

	flag.Parse()

	return Host, User, Password
	
}


