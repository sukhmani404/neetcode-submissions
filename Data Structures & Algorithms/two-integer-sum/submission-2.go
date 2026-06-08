func twoSum(nums []int, target int) []int {
    hash := map[int]int{}

	for i, num := range nums {
		if index, ok := hash[target-num]; ok {
			return []int{index, i}
		}

		hash[num] = i
	}

	return nil
}
