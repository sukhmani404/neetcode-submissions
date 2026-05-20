func trap(height []int) int {
	//for each element we should their extreme highs on either side
	//left side greatest for that element
	//and right side greatest for that element

	if len(height) == 0 { return 0 }
	nge := NGE(height)
	pge := PGE(height)
	width := 1
	total := 0
	
	for i, h := range height {
		if h < nge[i] && h < pge[i] {
			total += ((min(nge[i], pge[i]) - h)*width)
		}
	}

	return total
}

func NGE(height []int) []int {
	n := len(height)
	nge := make([]int, n)
	nge[n-1] = height[n-1]

	for i := n-2; i >=0; i-- {
		nge[i] = max(nge[i+1], height[i])
	}

	return nge
}

func PGE(height []int) []int {
	n := len(height)
	pge := make([]int, n)
	pge[0] = height[0]

	for i := 1; i < n; i++ {
		pge[i] = max(pge[i-1], height[i])
	}

	return pge
}
