func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charToCount := make(map[rune]int)

	for _, char := range s {
		charToCount[char]++
	}

	for _, char := range t {
		if _, ok := charToCount[char]; ok {
			charToCount[char]--
			if charToCount[char] == 0 {
				delete(charToCount, char)
			}
		}
	}

	if len(charToCount) > 0 {
		return false
	}

	return true
}