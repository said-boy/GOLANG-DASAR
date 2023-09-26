package helpers

func Variadic(number ...int) int {
	sum := 0
	for _, n := range number {
        sum += n
    }
	return sum
}

func VariadicWithData(number ...int) int {
	sum :=0 
	for _, n := range number {
		sum += n
	}
	return sum
}
