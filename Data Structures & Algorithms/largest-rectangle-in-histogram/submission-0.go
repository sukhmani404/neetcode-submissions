func largestRectangleArea(heights []int) int {

	pse := findPSE(heights)
	nse := findNSE(heights)

	maxArea := 0

	for i := 0; i < len(heights); i++ {
		maxArea = max(maxArea, (nse[i]-pse[i]-1)*heights[i])
	}

	return maxArea
}

func findPSE(arr []int) []int {

	pse := make([]int, len(arr))

	stack := []int{}

	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && arr[i] <= arr[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			pse[i] = stack[len(stack)-1]
		} else {
			pse[i] = -1

		}

		stack = append(stack, i)

	}

	return pse
}

func findNSE(arr []int) []int {

	nse := make([]int, len(arr))
	stack := []int{}

	for i := len(arr) - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[i] < arr[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			nse[i] = stack[len(stack)-1]
		} else {
			nse[i] = len(arr)
		}

		stack = append(stack, i)
	}

	return nse

}
