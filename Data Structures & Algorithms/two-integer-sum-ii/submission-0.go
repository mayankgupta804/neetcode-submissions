func twoSum(numbers []int, target int) []int {
	if len(numbers) <= 1 {
		return []int{-1, -1}
	}

	for i, j := 0, len(numbers) - 1; i < j; {
		sum := numbers[i] + numbers[j] 
		if sum == target {
			return []int{i+1, j+1}
		} else if sum < target {
			i++
		} else {
			j--
		}
	}

	return []int{-1, -1}
}
