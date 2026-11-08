func longestConsecutive(nums []int) int {
	mp := map[int]bool{}
	for _, num := range nums {
		mp[num] = true
	}

	var maxLen int

	for num := range mp{
		if !mp[num-1]{
			currNum := num
			currLen := 1

			for mp[currNum+1]{
				currNum++
				currLen++
			}

			maxLen = max(maxLen, currLen)
		}
	}

	return maxLen
}
