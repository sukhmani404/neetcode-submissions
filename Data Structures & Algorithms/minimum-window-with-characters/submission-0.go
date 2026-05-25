func minWindow(s string, t string) string {
    if len(s) < len(t) {
		return "" 
	}

	if t == "" {
		return ""
	}

	countT := map[byte]int{} // count of chars of t
	for i := 0; i < len(t); i++ {
		countT[t[i]] ++
	}

	have, need := 0, len(countT)
	res := []int{-1, -1}
	resLen := math.MaxInt32

	window := map[byte]int{}

	l := 0
	for r := 0; r < len(s); r++ {
		c := s[r]
		window[c]++
		if countT[c] > 0 && countT[c] == window[c] {
			have++
		}

		for have == need {
			//if we can find even min len than our resLen, update
			if (r - l + 1) < resLen {
				res = []int{l, r}
				resLen = r - l + 1
			}

			//see if we can shrink even more
			shrinked := s[l]
			window[shrinked]--
			//check if the shrinked value was in t
			//if it was then check if in window its count fell as compared to countT
			//if so then we need one more such character, so we decrease "have"
			if countT[shrinked] > 0 && window[shrinked] < countT[shrinked] {
				have--
			}

			l++
		}
	}

	if res[0] == -1 {
		return ""
	}

	return s[res[0]:res[1]+1]

}
