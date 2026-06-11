func longestConsecutive(nums []int) int {
	seen := map[int]bool{}

	for _, num := range nums {
		seen[num] = true
	}

	maxCount := 0

	for num := range seen {
		if !seen[num-1]{
			//start of seq
			curr := num
			currSeq := 1

			for seen[curr+1] {
				curr++
				currSeq++
			}

			maxCount = max(maxCount, currSeq)
		}
	}

	return maxCount
}
