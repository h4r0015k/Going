package array

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	length := len(numbersToSum)
	nums := make([]int, length)

	for index, toSum := range numbersToSum {
		nums[index] = Sum(toSum)
	}

	return nums
}

func SumAllTails(slices ...[]int) (sum []int) {

	for _, toSum := range slices {
		if len(toSum) == 0 {
			sum = append(sum, 0)
		} else {
			sum = append(sum, Sum(toSum[1:]))

		}
	}

	return
}
