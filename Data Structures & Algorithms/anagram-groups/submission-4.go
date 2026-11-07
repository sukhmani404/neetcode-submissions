func groupAnagrams(strs []string) [][]string {
	//we make a key and group them together
	groups := [][]string{}
	groupMap := map[[26]int][]string{}

	for _, str := range strs {
		key := [26]int{}
		for i := 0; i < len(str); i++ {
			key[str[i] - 'a']++
		}

		groupMap[key] = append(groupMap[key], str)
	}

	//group result together
	for _, group := range groupMap {
		groups = append(groups, group)
	}

	return groups
}
