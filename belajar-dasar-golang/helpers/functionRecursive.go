package helpers

func Fibonancy(number int) int {
	if number < 1 {
		return 1
	} else {
		return number * Fibonancy((number - 1))
	}
}	