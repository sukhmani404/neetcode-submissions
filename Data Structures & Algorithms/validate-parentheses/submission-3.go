func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false //if not pairs
	}

    opp := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}

	stack := []byte{}

	for i := 0; i < len(s); i++ {
		if startingBracket, closingBracket := opp[s[i]]; closingBracket {
			if len(stack) == 0 || stack[len(stack)-1] != startingBracket {
				return false
			}

			stack = stack[:len(stack)-1]
		}else{
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
