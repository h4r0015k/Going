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
