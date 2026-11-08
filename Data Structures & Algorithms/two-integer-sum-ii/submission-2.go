func twoSum(numbers []int, target int) []int {
	indexMap := map[int]int{}

	for i, num := range numbers {
		if ind, exists := indexMap[target - num]; exists {
			return []int{ind+1, i+1}
		}

		indexMap[num] = i
	}

	return nil
}
