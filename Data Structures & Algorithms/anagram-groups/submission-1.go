func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, str := range strs {
		//create a unique key
		key := [26]int{}
		for i := 0; i < len(str); i++ {
			key[str[i] - 'a']++
		}

		groups[key] = append(groups[key], str)
	}

	result := [][]string{}
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}
