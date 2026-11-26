func twoSum(nums []int, target int) []int {
  numToIdx := make(map[int]int)

  for idx, num := range nums {
	if targetIdx, ok := numToIdx[target-num]; ok {
		return []int{targetIdx, idx}
	}
	numToIdx[num] = idx
  }
  return []int{-1, -1}
}
