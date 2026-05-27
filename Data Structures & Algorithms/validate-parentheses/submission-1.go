func isValid(s string) bool {
    opp := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}

	stack := []byte{}

	for i := 0; i < len(s); i++ {
		if isOpeningBracket(s[i]) {
			stack = append(stack, s[i])
		}else if isClosingBracket(s[i]) {
			if len(stack) == 0 || opp[s[i]] != stack[len(stack)-1] {
				return false
			}

			stack = stack[:len(stack)-1]
		}	
		
	}

	return len(stack) == 0
}

func isOpeningBracket(b byte) bool {
	return b == '(' || b == '{' || b == '['
}

func isClosingBracket(b byte) bool {
	return b == ')' || b == '}' || b == ']'
}
