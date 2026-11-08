func productExceptSelf(nums []int) []int {
	left := 1
	right := 1

	ans := make([]int , len(nums))

	for i, num := range nums {
		ans[i] = left
		left *= num
	}

	for i := len(nums)-1; i >=0; i-- {
		ans[i] *= right
		right *= nums[i]
	}

	return ans
}
