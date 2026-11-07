func topKFrequent(nums []int, k int) []int {
	//we can use bucket sort for this

	if len(nums) == 0 {
		return nil //nothing to return
	}

	//get maxfreq and populate freq map
	maxFreq := 0
	freqMap := map[int]int{}

	for _, num := range nums {
		freqMap[num]++
		maxFreq = max(maxFreq, freqMap[num])
	}

	buckets := make([][]int, maxFreq+1) //avoid index panic while accessing

	//populate buckets
	for num, freq := range freqMap {
		buckets[freq] = append(buckets[freq], num)
	}


	ans := []int{}

	//take top k freq
	for i := maxFreq; i > 0; i-- {

		if len(buckets[i]) > 0 {
			ans = append(ans, buckets[i]...)
		}

		//if length exceeds k then break
		if len(ans) > k {
			ans = ans[:k]
			break
		}
	}

	return ans
}
