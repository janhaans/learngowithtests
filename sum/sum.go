package sum

func Sum(n []int) int {
	total := 0
	for _, v := range n {
		total += v
	}
	return total
}

func SumAll(slices ...[]int) []int {
	sums := []int{}

	for _, slice := range slices {
		sums = append(sums, Sum(slice))
	}

	return sums
}

func SumAllTails(slices ...[]int) []int {
	sums := []int{}

	for _, slice := range slices {
		if len(slice) <= 1 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(slice[1:]))
		}
	}

	return sums
}
