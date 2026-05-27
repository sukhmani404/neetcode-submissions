func evalRPN(tokens []string) int {
	stack := []int{}
	//keep pushing numbers on the stack
	//for each operation pop last two numbers on stack
	for _, token := range tokens {
		switch token {
			case "+", "-", "*", "/":
				b := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				a := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				var result int

				switch token {
					case "+" :
						result = a + b
					case "-":
						result = a - b
					case "*":
						result = a * b
					case "/":
						result = a / b 
				}
				
				stack = append(stack, result)
			
			default:
				num, _ := strconv.Atoi(token)
				stack = append(stack, num)
		}
	}

	return stack[len(stack)-1]
}

