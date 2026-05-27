func evalRPN(tokens []string) int {
	stack := []int{}
	//keep pushing numbers on the stack
	//for each operation pop last two numbers on stack
	for _, token := range tokens {
		var firstItem, secondItem int
		switch (token) {
			case "+":
				secondItem = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				firstItem = stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				stack = append(stack, firstItem + secondItem)

				
			case "-":
				secondItem = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				firstItem = stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				stack = append(stack, firstItem - secondItem)
			case "*":
				secondItem = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				firstItem = stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				stack = append(stack, firstItem * secondItem)
				
			case "/":
				secondItem = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				firstItem = stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				stack = append(stack, firstItem / secondItem)
				
			default:
				num, _ := strconv.Atoi(token)
				stack = append(stack, num)
		}
	}

	return stack[len(stack)-1]
}

