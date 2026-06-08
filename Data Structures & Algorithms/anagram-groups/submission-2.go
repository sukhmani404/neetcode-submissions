func groupAnagrams(strs []string) [][]string {
	ans := [][]string{}
	collect := map[string][]string{}

	for _, str := range strs {
		//form a key
		runes := []rune(str)
		sort.Slice(runes, func(i, j int)bool{
			return runes[i] < runes[j]
		})

		collect[string(runes)] = append(collect[string(runes)], str)
	}

	for _, v := range collect {
		ans = append(ans, v)
	}

	return ans
}
