func canConstruct(s string, k int) bool {
    if len(s) < k { return false }

    charCount := map[rune]int{}
    for _, rn := range s {
        charCount[rn]++
    }

    oddCounts := 0
    for _, v := range charCount {
        if v % 2 != 0 { oddCounts++ }
    }

    return oddCounts <= k
}