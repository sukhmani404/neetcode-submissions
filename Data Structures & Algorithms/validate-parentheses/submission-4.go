func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false //if not pairs
	}
	
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		char := s[i]
		switch (char) {
			case '(','{', '[':
				stack = append(stack, char)
			case ')':
				if len(stack) == 0 || stack[len(stack)-1] != '('{
					return false
				}
				stack = stack[:len(stack)-1]
			case '}':
				if len(stack) == 0 || stack[len(stack)-1] != '{'{
					return false
				}
				stack = stack[:len(stack)-1]
			
			case ']':
				if len(stack) == 0 || stack[len(stack)-1] != '['{
					return false
				}
				stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
