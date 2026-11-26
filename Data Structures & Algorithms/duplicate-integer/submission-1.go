func hasDuplicate(nums []int) bool {
	numToCount := make(map[int]int)

	for _, num := range nums {
		numToCount[num]++
		if numToCount[num] > 1 {
			return true
		}
	}

	return false
}
