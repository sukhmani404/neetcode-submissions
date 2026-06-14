func maxArea(heights []int) int {

	maxarea := 0

	i, j := 0, len(heights)-1

	for i < j {
		height := min(heights[i], heights[j])
		width := j - i
		area := width * height
		maxarea = max(maxarea, area)
		if heights[i] < heights[j] {
			i++
		}else {
			j--
		}
	}

	return maxarea
}
