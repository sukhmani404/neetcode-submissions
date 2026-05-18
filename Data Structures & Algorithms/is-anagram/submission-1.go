func isAnagram(s string, t string) bool {
	//increase count of character in one pass 
	//decrease in another 
	//if we encounter any negative freq during decreasing the count
	//strings are not anagram, some other character is present here

	//check for freq in another pass

	freq := map[rune]int{}
	for _, r := range s {
		freq[r] ++
	}

	for _, r := range t {
		freq[r]--
	}

	for _, f := range freq {
		if f != 0 {
			return false
		}
	}
	
	return true
}
