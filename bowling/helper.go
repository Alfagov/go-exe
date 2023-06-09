package bowling

// sum a list of ints helper function
func sum(input ...int) int {
	sum := 0
	for _, i := range input {
		sum += i
	}
	return sum
}
