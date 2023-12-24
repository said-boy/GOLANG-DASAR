package helpers

type Filter func(int) string

func SayStatus(umur int, GetStatus Filter ) string {
	result := GetStatus(umur)
	return "Kamu tegolong : " + result
}

func GetStatus(umur int) string {
	if umur < 18 {
		return "Muda"
    } else {
		return "Tua"
	}
}