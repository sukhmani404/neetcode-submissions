func groupAnagrams(strs []string) [][]string {
	//sort all the strings and make a key and append those keys together and then form a response?
	hashmp := map[string][]string{}

	for _, str := range strs {
		s := []rune(str)
		sort.Slice(s, func(i, j int)bool{
			return s[i] < s[j]
		})
		
		hashmp[string(s)] = append(hashmp[string(s)], str)
	}

	result := [][]string{}

	for _, group := range hashmp {
		result = append(result, group)
	}

	return result
}
