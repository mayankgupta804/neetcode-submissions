func topKFrequent(nums []int, k int) []int {
	numToCount := make(map[int]int)
	for _, v := range nums {
		numToCount[v] += 1
	}

	countToNumBuckets := make([][]int, len(nums) + 1)

	for k, v := range numToCount {
		if countToNumBuckets[v] == nil {
			countToNumBuckets[v] = make([]int,0)
		}
		countToNumBuckets[v] = append(countToNumBuckets[v], k)
	}
	
	topKCount := k
	out := make([]int, 0)

	for i:=len(countToNumBuckets)-1; i>=0; i-- {
		if topKCount == 0 {
			return out
		}
		if len(countToNumBuckets[i]) > 0 && topKCount > 0 {
			for j:=0; j<len(countToNumBuckets[i]); j++ {
				out = append(out, countToNumBuckets[i][j])
				topKCount -= 1
			}
		}
	}

	return out
}
