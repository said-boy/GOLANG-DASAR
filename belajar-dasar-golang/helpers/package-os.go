package helpers

import "os"

// untuk dapat berinteraksi dengan komputer kita

func PrintHostname() (name string, err error) {
	return os.Hostname()
}

