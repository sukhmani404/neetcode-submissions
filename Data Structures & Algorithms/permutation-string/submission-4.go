func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) { return false }
    //use sliding window of s1 size over s2

    window := [26]int{} //only increase or decrease bytes that are common
    charMap := map[byte]bool{}
    for i := 0; i < len(s1); i++ {
        window[s1[i] - 'a']++ 
        charMap[s1[i]] = true
    }

    windowSize := len(s1)
    for i := 0; i + windowSize - 1 < len(s2); i++ {
        windowToCompare := [26]int{}
        substr := s2[i:i+windowSize]
        fmt.Println(substr)
        //then loop over substr
        for j := 0; j < len(substr); j++ {
            if !charMap[substr[j]] {
                break //if this character is not in s1
            }

            windowToCompare[substr[j] - 'a']++
        }

        if window ==  windowToCompare {
            return true
        }
    }
    return false
}