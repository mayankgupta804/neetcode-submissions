import "slices"

func groupAnagrams(strs []string) [][]string {
	strToAnagrams := make(map[string][]string)

	for _, str := range strs {
		runes := []rune(str)
		slices.Sort(runes)
		sortedStr := string(runes)

		if _, ok := strToAnagrams[sortedStr]; !ok {
			strToAnagrams[sortedStr] = make([]string, 0)
		}
		strToAnagrams[sortedStr] = append(strToAnagrams[sortedStr], str)
	}

	anagramGroups := make([][]string, 0)

	for _, groups := range strToAnagrams {
		anagramGroups = append(anagramGroups, groups)
	}
	
	return anagramGroups
}

