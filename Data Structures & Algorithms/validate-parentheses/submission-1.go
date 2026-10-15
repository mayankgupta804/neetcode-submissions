func isValid(s string) bool {
    if len(s) == 1 {
		return false
	}

	bMap := map[rune]rune{
		')':'(',
		'}':'{',
		']':'[',
	}

	stack := make([]rune,0)

	for _, br := range s {
		if len(stack) > 0 && stack[len(stack) - 1] == bMap[br] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, br)
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}
