func isPalindrome(s string) bool {
	i, j := 0, len(s) - 1

	for i<j {
		for i<j && !isAlphaNumeric(rune(s[i])) {
			i++
		}

		for i<j && !isAlphaNumeric(rune(s[j])) {
			j--
		}

		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}

		i++
		j--
	}

	return true
}

func isAlphaNumeric(char rune) bool {
	if (char >= '0' && char <='9') || (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
		return true
	}
	return false
}