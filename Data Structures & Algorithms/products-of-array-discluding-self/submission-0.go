func productExceptSelf(nums []int) []int {
	//keep track of the product 
	//solve it in two passes

	leftSideProduct := 1
	result := make([]int, len(nums))

	//here we skip current element for each index
	for i, num := range nums {
		result[i] = leftSideProduct
		leftSideProduct *= num
	}

	rightSideProduct := 1
	//multiply each of the elements in the array with result
	//except current. same strat
	for i := len(nums)-1; i >= 0; i-- {
		result[i] *= rightSideProduct
		rightSideProduct *= nums[i]
	}

	return result
}
