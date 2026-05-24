func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) { 
        return false 
    }

    window := [26]int{} 
    target := [26]int{}

    for i := 0; i < len(s1); i++ {
        target[s1[i] - 'a']++
        window[s2[i] - 'a']++
    }

    if target == window {
        return true
    }

    left := 0

    for right := len(s1); right < len(s2); right++ {
        window[s2[right]- 'a']++  //add right
        window[s2[left] - 'a']-- //remove left
        left++ //move boundary

        if target == window {
            return true
        }
    }

    return false
    
}