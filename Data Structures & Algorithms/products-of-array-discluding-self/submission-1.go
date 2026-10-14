func productExceptSelf(nums []int) []int {
	prefixProd := make([]int, len(nums))

	for i, _ := range nums{
		if i == 0 {
			prefixProd[i] = 1
		} else if i == 1 {
			prefixProd[i] = nums[i-1]
		} else {
			prefixProd[i] = prefixProd[i-1] * nums[i-1]
		}
	}

	suffixProd := make([]int, len(nums))

	for i:=len(nums)-1; i>=0; i-- {
		if i == len(nums) - 1 {
			suffixProd[i] = 1
		} else if i == len(nums) - 2 {
			suffixProd[i] = nums[len(nums)-1]
		} else {
			suffixProd[i] = suffixProd[i+1] * nums[i+1]
		}
	}

	for i:=0; i<len(nums); i++ {
		nums[i] = prefixProd[i] * suffixProd[i]
	}

	return nums
}

// 2 3 5 6 

// 1 2 6 30

// 90 30 6 1

// 90 60 36 30

// 1 2 4 6

// 1 1 2 8
// 48 24 6 1 

// 48 24 12 8