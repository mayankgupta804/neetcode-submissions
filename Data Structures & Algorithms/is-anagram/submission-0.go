func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    charToCount := make(map[rune]int, len(s))

    for _, v := range s {
        charToCount[v] += 1
    }

    for _, v := range t {
        if _, ok := charToCount[v]; ok {
            charToCount[v] -= 1
        }
    }

    for _, v := range charToCount {
        if v > 0 {
            return false
        }
    }

    return true
}
