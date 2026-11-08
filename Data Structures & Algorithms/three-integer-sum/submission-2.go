func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	ans := [][]int{}

	for i := 0; i < len(nums) - 2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i+1
		k := len(nums)-1

		for j < k {

		sum := nums[i] + nums[j] + nums[k]

		if sum == 0 {
			ans = append(ans, []int{nums[i], nums[j], nums[k]})
			//skip duplicates
			j++
			for j < k && nums[j] == nums[j-1]{
				j++
			}

			k--
			for j < k && nums[k] == nums[k+1] {
				k--
			}
		}else if sum > 0 {
			k--
		}else {
			j++
		}

		}
	}

	return ans
}
