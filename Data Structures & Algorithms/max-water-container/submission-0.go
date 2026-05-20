func maxArea(heights []int) int {
	i, j := 0, len(heights)-1

	maxarea := 0
	for i < j {
		width := j-i
		height := min(heights[i], heights[j])
		area := height * width
		maxarea = max(maxarea, area)
		if heights[i] > heights[j] {
			j--
		}else {
			i++
		}
	}

	return maxarea
}
