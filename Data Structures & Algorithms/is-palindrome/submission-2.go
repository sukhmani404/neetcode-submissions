func isPalindrome(s string) bool {
	i := 0
	j := len(s)-1

	for i < j {
		a, ok := isAlphaNumeric(s[i])
		if !ok {
			i++
			continue
		}
	
		b, ok := isAlphaNumeric(s[j])
		if !ok {
			j--
			continue
		}

		if a != b {
			return false
		}

		i++
		j--
	}

	return true
}

func isAlphaNumeric(b byte) (byte, bool){
	if b >= 48 && b <= 57 {
		return b, true
	}

	if b >= 97 && b <= 122 {
		return b, true
	}

	if b >= 65 && b <= 90 {
		return b + 32, true
	}

	return 0, false
}
