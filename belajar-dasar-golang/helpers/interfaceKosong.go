package helpers

// interface kosong memungkinkan anda dapat memberikan
// parameter bertipe bebas. 

func Message(t ...interface{}) interface{} {
	return t
}


