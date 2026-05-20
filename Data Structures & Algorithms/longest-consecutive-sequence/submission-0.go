func longestConsecutive(nums []int) int {
	//use map to store all the nums first
	mp := map[int]struct{}{}
	for _, num := range nums {
		mp[num] = struct{}{}
	}

	maxLen := 0
	for key := range mp {
		if _, ok := mp[key-1]; !ok {
			//start of the sequence
			currLen := 0
			currNum := key

			for {
				if _, ok := mp[currNum]; !ok {
					break
				}

				currNum++
				currLen++
			}

			maxLen = max(maxLen, currLen)
		}
	}

	return maxLen
}
