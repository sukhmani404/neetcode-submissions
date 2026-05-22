func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	left := 0
	lastSeen := map[byte]int{}
	maxlen := 0
	for right := 0; right < len(s); right ++{
		if index, seen := lastSeen[s[right]]; seen && index >= left {
			left = index+1
		}

		lastSeen[s[right]] = right 

		currlen := right - left + 1
		maxlen = max(maxlen, currlen)

	}

	return maxlen
}
