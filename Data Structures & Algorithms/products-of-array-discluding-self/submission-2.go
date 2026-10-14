func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))

	for i, _ := range nums{
		if i == 0 {
			answer[i] = 1
		} else if i == 1 {
			answer[i] = nums[i-1]
		} else {
			answer[i] = answer[i-1] * nums[i-1]
		}
	}

	suffixProd := 1

	for i:=len(answer)-1; i>=0; i-- {
		answer[i] = answer[i] * suffixProd
		suffixProd *= nums[i]
	}

	return answer
}