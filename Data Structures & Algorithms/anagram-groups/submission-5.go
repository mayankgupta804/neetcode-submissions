import "slices"

func groupAnagrams(strs []string) [][]string {
	charsToAnagrams := make(map[string][]string)
	for _, str := range strs {
		s := sortString(str)
		charsToAnagrams[s] = append(charsToAnagrams[s], str)
	}

	answer := make([][]string, 0)
	for _, strArr := range charsToAnagrams {
		answer = append(answer, strArr)
	}
	
	return answer
}

func sortString(s string) string {
	if strings.TrimSpace(s) == "" {
		return s
	}
	runes := []rune(s)
	slices.Sort(runes)
	return string(runes)
}
