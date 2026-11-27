func productExceptSelf(nums []int) []int {
	leftProduct := make([]int,0)

	for i:=0; i<len(nums); i++ {
		if i == 0 {
			leftProduct = append(leftProduct, 1)
		} else if i == 1 {
			leftProduct = append(leftProduct, nums[i-1])
		} else {
			leftProduct = append(leftProduct, nums[i-1] * leftProduct[i-1])
		}
	}

	rightProduct := make([]int, len(nums))

	for i:=len(nums)-1; i>=0; i-- {
		if i == len(nums)-1 {
			rightProduct[i] = 1
		} else if i == len(nums) - 2 {
			rightProduct[i] = nums[len(nums)-1]
		} else {
			rightProduct[i] = rightProduct[i+1] * nums[i+1]
		}
	}

	for i:=0; i<len(nums); i++ {
		nums[i] = leftProduct[i]*rightProduct[i]
	}

	return nums
}

// 1 1 2 8

// 48  24  6  1

// 48 24 12 8