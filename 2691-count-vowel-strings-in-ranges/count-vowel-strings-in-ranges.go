func vowelStrings(words []string, queries [][]int) []int {
    vowels := map[byte]bool{ 'a': true, 'e': true, 'i': true, 'o': true, 'u': true }
    before := make([]int, len(words))
    isVowel := make([]int, len(words))
    results := make([]int, len(queries))

    for i, word := range words {
        if i > 0 { 
            before[i] = before[i-1]
            isVowel[i] = 0
        }

        _, okStart := vowels[word[0]]
        _, okEnd := vowels[word[len(word)-1]]

        if okStart && okEnd {
            before[i] += 1
            isVowel[i] = 1
        }
    }

    for i, query := range queries {
        results[i] = before[query[1]] - before[query[0]] + isVowel[query[0]]
    }

    return results
}