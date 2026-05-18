func twoSum(nums []int, target int) []int {
    if len(nums) < 2 {
		return nil
	}

	if len(nums) == 2 {
		return []int{0, 1}
	}

	mpIndex := map[int]int{}

	for i, num := range nums {
		if index, found := mpIndex[target - num]; found {
			return []int{min(index, i), max(index, i)}
		}

		mpIndex[num] = i
	}

	return nil //impossible condition
}
