func hasDuplicate(nums []int) bool {
    hashMap := map[int]struct{}{}

	for _, num := range nums {
		if _, exists := hashMap[num]; exists {
			return true
		}

		hashMap[num] = struct{}{}
	}

	return false
}
