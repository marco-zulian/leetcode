func strStr(haystack string, needle string) int {
    for i := range(haystack) {
        if len(haystack) - i < len(needle) { break }

        idx := i
        for haystack[idx] == needle[idx-i] {
            idx++
            if (idx - i) == len(needle) {
                return i
            }
        }
    }

    return -1
}