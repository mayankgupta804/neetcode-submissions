func isPalindrome(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}

	for i,j := 0, len(s)-1; i<j;{
		for i<j && !isASCIIAlphanumeric(s[i]) {
			i++
		}

		for j>i && !isASCIIAlphanumeric(s[j]) {
			j--
		}

		if !strings.EqualFold(string(s[i]), string(s[j])) {
			return false
		}

		i+=1
		j-=1 
	}
	return true
}

func isASCIIAlphanumeric(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9')
}
