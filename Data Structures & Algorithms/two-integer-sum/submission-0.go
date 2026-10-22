func twoSum(nums []int, target int) []int {
    numToIdx := make(map[int]int)
	for currIdx, num := range nums {
		if idx, ok := numToIdx[target-num]; ok {
			return []int{idx,currIdx}
		}
		numToIdx[num] = currIdx
	}
	return []int{-1, -1}
}
