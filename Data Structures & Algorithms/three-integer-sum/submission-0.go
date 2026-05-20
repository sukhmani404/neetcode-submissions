func threeSum(nums []int) [][]int {
	//same two pointer approach can be done here
	if len(nums) == 0 {
		return nil
	}

	sort.Ints(nums) //sort so that we can apply two pointer approach

	result := [][]int{}

	n := len(nums)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}


		j, k := i+1, n-1
		//fix i and change j and k and try to find triplets
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				//save result
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--

				//skip duplicates if any
				for j < k && nums[j] == nums[j-1] {
					j++
				}

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

	return result
}
