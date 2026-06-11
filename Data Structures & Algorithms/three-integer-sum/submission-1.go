func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	ans := [][]int{}

	for i, num := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j, k := i+1, len(nums)-1
		for j < k {
		sum := num + nums[j] + nums[k]

			if sum == 0 {
				ans = append(ans, []int{num, nums[j], nums[k]})
				//skip duplicates
				for j < k && nums[j] == nums[j+1] {
					j++
				}

				j++

				for j < k && nums[k] == nums[k-1] {
					k--
				}
			}else if sum > 0 {
				k--
			}else{
				j++
			}
		}
	}

	return ans
}
