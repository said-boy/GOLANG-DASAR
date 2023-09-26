package helpers

func IsAdmin(user string, isAdmin func(string) bool) string{
	if isAdmin(user) {
		return "Selamat datang " + user
	} else {
		return "User " + user + " ditolak, Halaman hanya untuk Admin"
	}
}