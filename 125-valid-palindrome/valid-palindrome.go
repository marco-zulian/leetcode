func isPalindrome(s string) bool {
    lowercased := strings.ToLower(s)
    alphanumRegex, _ := regexp.Compile("[^0-9a-z]+")
    trimmed := alphanumRegex.ReplaceAllString(lowercased, "")
    left, right := 0, len(trimmed) - 1

    for left < right {
        if trimmed[left] != trimmed[right] {
            return false
        }

        left++
        right--
    }

    return true
}

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}