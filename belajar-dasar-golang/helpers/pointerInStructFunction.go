package helpers

type Akun struct {
	Username, Password string
}

// untuk dapat merubah aslinya, gunakan pointer
func ChangePassword(user *Akun){
	user.Password = "SecurePassword"
}

func ChangePasswordNoPointer(user Akun){
	user.Password = "SecurePassword"
}
