func productExceptSelf(nums []int) []int {
	left := 1
	ans := make([]int, len(nums))
	for i, num := range nums {
		ans[i] = left
		left *= num
	}

	right := 1

	for i := len(nums)-1; i>=0; i-- {
		num := nums[i]

		ans[i] *=  right
		right *= num
	}

	return ans
}
