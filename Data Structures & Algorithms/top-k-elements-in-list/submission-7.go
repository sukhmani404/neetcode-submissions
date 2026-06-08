func topKFrequent(nums []int, k int) []int {
	freq := map[int]int{}
	maxFreq := 0

	for _, num := range nums {
		freq[num]++
		maxFreq = max(maxFreq, freq[num])
	}

	buckets := make([][]int, maxFreq + 1)

	for c, f := range freq {
		buckets[f] = append(buckets[f], c)
	}

	ans := []int{}

	for f := maxFreq; f >= 0; f-- {
		ans = append(ans, buckets[f]...)

		if len(ans) > k {
			ans = ans[:k]
			break
		}
	}

	return ans
}
