func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	
	var b strings.Builder

	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9') || (s[i] >= 'A' && s[i] <= 'Z'){
			b.WriteByte(s[i])
		}
	}

	c := b.String()
	i := 0
	j := len(c)-1

	for i < j {
		if c[i] != c[j] {
			return false
		}

		i++
		j--
	}

	return true
}
