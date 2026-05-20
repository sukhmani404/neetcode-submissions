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
		if diff == 1 {
			currLen++
		}else if diff != 0 {
			currLen = 1
		}
		currNum = num
		maxLen = max(maxLen, currLen)
	}

	return maxLen
}
