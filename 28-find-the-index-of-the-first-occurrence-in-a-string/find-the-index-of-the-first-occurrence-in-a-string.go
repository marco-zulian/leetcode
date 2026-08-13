func strStr(haystack string, needle string) int {
    needle_rn := []rune(needle)
    haystack_rn := []rune(haystack)
    for i, rn_h := range haystack {
        if i + len(needle) > len(haystack) {
            break
        }
        if rn_h == needle_rn[0] {
            k := i
            for _, rn_n := range needle {
                if haystack_rn[k] == rn_n {
                    k++
                    continue
                }
                break
            }
            if (k-i) == len(needle) {
                return i
            }
        }
    }
    return -1
}