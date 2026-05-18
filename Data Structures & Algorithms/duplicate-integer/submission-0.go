func hasDuplicate(nums []int) bool {
    //simple stuff is to just use a hash map
    //and check when updating each entry that whether the frequency is greater than one or not

    hashmap := map[int]int{}
    for _, num := range nums {
        hashmap[num]++
        if hashmap[num] >1 {
            return true
        }
    }

    return false
}
