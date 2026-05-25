func characterReplacement(s string, k int) int {
	freq := map[byte]int{}
	res, l, f := 0, 0, 0
	
	for r := 0; r < len(s); r ++ {
		freq[s[r]]++ //increase freq
		f = max(f, freq[s[r]]) //compare max freq

		//shrink window if non frequent characters are more than k in the window
		for (r - l + 1) - f > k {
			freq[s[l]]--
			l++
		}

		res = max(res, r - l + 1)
	}

	return res
}

