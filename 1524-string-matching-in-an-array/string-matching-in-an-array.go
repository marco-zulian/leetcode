func stringMatching(words []string) []string {
    substrings := map[string]bool{}
    result := []string{}

    for i, word := range words {
        if _, ok := substrings[word]; ok { continue }

        for _, otherWord := range words[i+1:] {
            if strings.Contains(word, otherWord) {
                substrings[otherWord] = true
            }

            if strings.Contains(otherWord, word) {
                substrings[word] = true
                break
            }
        }
    }

    for k, _ := range substrings {
        result = append(result, k)
    }

    return result
}