func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashMap := map[rune]int{}

	for _, r := range s {
		hashMap[r]++
	}

	for _, r := range t {
		hashMap[r]--
		if hashMap[r] < 0 {
			return false
		}
	}

	return true
}
