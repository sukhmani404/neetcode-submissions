func twoSum(nums []int, target int) []int {
    indexMap := map[int]int{}

	for i, num := range nums {
		if ind, exists := indexMap[target - num]; exists {
			return []int{ind, i}
		}

		indexMap[num] = i
	}

	return nil //impossible to reach given question requirements
}
