func groupAnagrams(strs []string) [][]string {
	ans := [][]string{}
	collect := map[[26]int][]string{}

	for _, str := range strs {
		key := [26]int{}
		for i := 0; i < len(str); i++ {
			key[str[i] - 'a']++
		}

		collect[key] = append(collect[key], str)
	}

	for _, v := range collect {
		ans = append(ans, v)
	}

	return ans
}
