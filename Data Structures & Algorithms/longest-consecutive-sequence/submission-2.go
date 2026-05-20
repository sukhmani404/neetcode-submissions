func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)
	//2, 20, 4, 10, 3, 4, 5
	//2, 3, 4, 5, 10, 20

	currLen := 1
	maxLen := 1
	currNum := nums[0]
	
	for _, num := range nums[1:] {
		diff := num - currNum
		switch diff {
			case 0:
			case 1:
				currLen++
				currNum = num
			default:
				currLen = 1
				currNum = num
		}

		maxLen = max(maxLen, currLen)
	}

	return maxLen
}
