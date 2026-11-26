import "slices"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s1 := sortString(s)
	t1 := sortString(t)

	return s1 == t1
}

func sortString(s string) string {
	runes := []rune(s)
	slices.Sort(runes)
	return string(runes)
}