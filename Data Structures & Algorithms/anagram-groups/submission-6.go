func groupAnagrams(strs []string) [][]string {
	arrToAna := make(map[[26]int][]string)

	for _, str := range strs {
		arr := [26]int{}
		for _, char := range str {
			arr[char-97]++
		}
		arrToAna[arr] = append(arrToAna[arr], str)
	}

	answer := make([][]string, 0)

	for _, strArr := range arrToAna {
		answer = append(answer, strArr)
	} 

	return answer
}
