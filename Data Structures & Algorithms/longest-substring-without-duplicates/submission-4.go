func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
        return 0
    }

    maxlen := 0
    l := 0
    seen := map[byte]int{}

    for r := 0; r < len(s); r++ {
        if index, ok := seen[s[r]]; ok && index >= l {
            l = index + 1 //move left boundary
        }

        seen[s[r]] = r

        currlen := r - l + 1
        maxlen = max(maxlen, currlen)
    }

    return maxlen
}
