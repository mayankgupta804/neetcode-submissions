func longestConsecutive(nums []int) int {
	numToExists := make(map[int]struct{})
	// create a set
	for _, v := range nums {
		numToExists[v] = struct{}{}
	}

	var longest = 0

	for _, num := range nums {
		if _, ok := numToExists[num-1]; !ok {
			length := 0
			for _, exists := numToExists[num+length]; exists; _, exists = numToExists[num+length] {
				length += 1
			}
			if length > longest {
				longest = length
			}
		}
	}

	return longest
}
