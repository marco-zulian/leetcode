func wordSubsets(words1 []string, words2 []string) []string {
    requiredLetters := map[rune]int{}
    result := []string{}

    for _, word := range words2 {
        requiredLettersForWord := map[rune]int{}

        for _, rn := range word {
            requiredLettersForWord[rn]++

            if v, ok := requiredLetters[rn]; !ok || requiredLettersForWord[rn] > v {
                requiredLetters[rn] = requiredLettersForWord[rn]
            }
        }
    }

    for _, word := range words1 {
        availableLettersOnWord := map[rune]int{}

        for _, rn := range word {
            availableLettersOnWord[rn]++
        }

        isUniversal := true
        for rn,needed := range requiredLetters {
            if available, ok := availableLettersOnWord[rn]; !ok || needed > available {
                isUniversal = false
                break
            }
        }

        if isUniversal {
            result = append(result, word)
        }
    }

    return result
}