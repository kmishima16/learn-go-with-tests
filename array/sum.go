package array

func Sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func SumAll(slices ...[]int) []int {
	var sums []int
	for _, slice := range slices {
		sums = append(sums, Sum(slice))
	}
	return sums
}