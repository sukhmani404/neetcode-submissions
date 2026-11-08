type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	//build string with len + special character + actual string
	var sb strings.Builder

	for _, str := range strs {
		sb.WriteString(strconv.Itoa(len(str)))
		sb.WriteByte('#')
		sb.WriteString(str)
	}

	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	//break the string down
	//find the first occurance of #, read the first part and get the len
	//traverse the string, read substr and then read until # to get the len again

	strs := []string{}

	i := 0

	for i < len(encoded) {
		//find the first occurance
		j := i
		
		for j < len(encoded) && encoded[j] != '#' {
			j++
		}
		

		length, _ := strconv.Atoi(encoded[i:j])
		str := encoded[j + 1 : j + 1 + length]
		strs = append(strs, str)
		i = j + 1 + length
	}

	return strs
}
