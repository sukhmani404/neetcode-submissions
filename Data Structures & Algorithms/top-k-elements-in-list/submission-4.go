func topKFrequent(nums []int, k int) []int {
	freqMap := map[int]int{}
	maxFreq := 0

	for _, num := range nums {
		freqMap[num]++
		maxFreq = max(maxFreq, freqMap[num])
	}

	buckets := make([][]int, maxFreq+1) //+1 cause maxFreq can also be an index

	for num, freq := range freqMap {
		buckets[freq] = append(buckets[freq], num)
	}

	result := []int{}
	//now select top k frequencies
	for f := maxFreq; f >= 0; f-- {
		bucket := buckets[f]
		result = append(result, bucket...)

		if len(result) >= k {
			result = result[:k]
			break
		}
	}

	return result
}