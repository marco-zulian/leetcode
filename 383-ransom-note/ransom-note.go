func canConstruct(ransomNote string, magazine string) bool {
    availableRunes := map[rune]int{}

    for _, rn := range magazine {
        availableRunes[rn]++
    }

    for _, rn := range ransomNote {
        if count, ok := availableRunes[rn]; !ok || count == 0 {
            return false
        }

        availableRunes[rn]--
    }

    return true
}