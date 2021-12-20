package mate

func Sum(numbers ...int) int {
	var sum = 0
	for _, i := range numbers {
		sum += i
	}
	return sum
}
