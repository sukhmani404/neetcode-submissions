func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	nge, pge := NGE(height), PGE(height)
	total := 0

	for i := 0; i < len(height); i++ {
		if height[i] < nge[i] && height[i] < pge[i] {
			total += (min(nge[i], pge[i]) - height[i])
		}
	}

	return total
}

func NGE(height []int) []int {
	n := len(height)
	nge := make([]int, n)

	nge[n-1] = height[n-1]
	for i := n-2; i >0; i-- {
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