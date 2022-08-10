package learn_ci_cd

func Sum(numbers ...int) int {
	var sum int
	for _, number := range numbers {
		sum += number + 1
	}
	return sum
}
