import "slices"

func hasDuplicate(nums []int) bool {
	slices.Sort(nums)

	for i, num := range nums {
		if i>0 && nums[i-1] == num {
			return true
		}
	}
	return false
}
