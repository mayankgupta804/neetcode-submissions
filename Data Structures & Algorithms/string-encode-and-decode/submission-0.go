type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var out string
	for _, str := range strs {
		out += strconv.Itoa(len(str)) + "#" + str
	}
	return out
}

func (s *Solution) Decode(encoded string) []string {

	res := make([]string, 0)

	for i := 0; i < len(encoded); {
		j := i

		for encoded[j] != '#' {
			j += 1
		}
		length, _ := strconv.Atoi(encoded[i:j])
		res = append(res, encoded[j+1:j+1+length])
		i = j + 1 + length
	}

	return res
}

