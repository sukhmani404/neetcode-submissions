func isPalindrome(s string) bool {
	r := extractAlphaNumStr(s)
	runes := []rune(r)
	reverse(runes, 0, len(runes)-1)
	return r == string(runes)
}

func reverse(runes []rune, start, end int) {
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}
}

func extractAlphaNumStr(s string) string {
	var str strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if smallAlph(b) || isNum(b) {
			str.WriteByte(b)
		}

		if bigAlph(b) {
			str.WriteByte(b + 32)
		}
	}

	return str.String()
}

func smallAlph(b byte) bool {
	return b >= 'a' && b <= 'z'  
}

func bigAlph(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func isNum(b byte) bool {
	return b >= '0' &&  b <= '9'
}
