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
	var res []string
	
	i := 0
	for i < len(encoded) {
		//find the separator that we added
		j := i
		for j < len(encoded) && encoded[j] != '#' {
			j++
		}

		//length is encoded[i:j]
		length, _ := strconv.Atoi(encoded[i:j])
		//skip '#'
		i = j + 1

		//read the string of given length
		str := encoded[i: i+ length]
		res = append(res, str)
		i += length
	}

	return res
}
