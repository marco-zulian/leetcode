func firstUniqChar(s string) int {
    letterCount := map[rune]int{}

    for _, rn := range s {
        letterCount[rn]++
    }

    for i, rn := range s {
        if count, _ := letterCount[rn]; count < 2 {
            return i
        }
    }

    return -1
}