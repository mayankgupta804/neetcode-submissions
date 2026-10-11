import "slices"

func hasDuplicate(nums []int) bool {
    if len(nums) == 1 {
		return false
	}

	slices.Sort(nums)
	var val int
	for i:= 0; i<len(nums); i++ {
		if i + 1 < len(nums) {
			val = nums[i]^nums[i+1]
			if val == 0 {
				return true
			}
		}
	}

	return false
}
