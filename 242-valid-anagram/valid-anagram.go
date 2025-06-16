func isAnagram(s string, t string) bool {
    if len(s) != len(t) { return false }

    letterCount := map[rune]int{}
    for _, rn := range s {
        letterCount[rn]++
    }

    for _, rn := range t {
        if val, ok := letterCount[rn]; val <= 0 || !ok {
            return false
        }

        letterCount[rn]--
    }

    return true
}