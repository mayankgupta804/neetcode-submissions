import "slices"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	slices.Sort(nums)

	answer := make([][]int, 0)

	for i, v := range nums{
		if i > 0 && v == nums[i-1] {
			continue
		}
		target := -nums[i]
		for j,k := i+1, len(nums) - 1; j<k; {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				answer = append(answer, []int{nums[i], nums[j], nums[k]})
				j++
				for nums[j] == nums[j-1] && j<k{
					j++
				} 
				k--
			} else if nums[j] + nums[k] > target {
				k--
			} else {
				j++
			}
		}
	}
	return answer
}
