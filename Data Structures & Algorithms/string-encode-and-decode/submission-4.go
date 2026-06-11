type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var b strings.Builder
	for _, str := range strs {
		b.WriteString(strconv.Itoa(len(str)))
		b.WriteByte('#')
		b.WriteString(str)
	}

	return b.String()
}

func (s *Solution) Decode(encoded string) []string {
	var ans []string
	i := 0

	for i < len(encoded) {
		j := i
		//find out the occurrance of our separator
		for j < len(encoded) && encoded[j] != '#' {
			j++
		}

		length, _ := strconv.Atoi(encoded[i:j])

		i = j + 1

		str := encoded[i: i + length]
		ans = append(ans, str)
		i += length
	}

	return ans
}
