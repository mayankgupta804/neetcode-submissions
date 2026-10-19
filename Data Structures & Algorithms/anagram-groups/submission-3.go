func groupAnagrams(strs []string) [][]string {
	countToStr := make(map[[26]int][]string)

	for _, str := range strs {
		countArr := [26]int{}
		for _, char := range str {
			countArr[int(char) - int('a')]++
		}
		if _, ok := countToStr[countArr]; !ok {
			countToStr[countArr] = make([]string, 0)
		}
		countToStr[countArr] = append(countToStr[countArr], str)
	}

	anagramGroups := make([][]string, 0)
	for _, groups := range countToStr {
		anagramGroups = append(anagramGroups, groups)
	}

	return anagramGroups
}

