func lengthOfLongestSubstring(s string) int {
    chars := map[rune]bool{}
    i, max := 0, 0

    for j, rn := range s {
        if seen, ok := chars[rn]; ok && seen {
            if j - i > max {
                max = j - i
            }

            for i < j {
                chars[rune(s[i])] = false
                if rune(s[i]) == rn {
                    i++
                    break
                }
                i++
            }
        }

        chars[rn] = true
    }

    if len(s) - i > max {
        max = len(s) - i
    }

    return max
}