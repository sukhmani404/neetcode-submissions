func maxArea(heights []int) int {
	start, end := 0, len(heights)-1
	mx := 0

	for start < end {
		currArea := min(heights[start], heights[end]) * (end - start)
		mx = max(mx, currArea)
		if heights[start] > heights[end] {
			end--
		}else {
			start++
		}
	}

	return mx
}
