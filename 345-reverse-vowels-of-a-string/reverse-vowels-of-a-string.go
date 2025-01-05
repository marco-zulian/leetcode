func reverseVowels(s string) string {
    vowelIndices := []int{}
    result := []rune(s)

    for i, rn := range s {
        if isVowel(rn) {
            vowelIndices = append(vowelIndices, i)
        }
    }

    i, j := 0, len(vowelIndices) - 1
    for i < j {
        result[vowelIndices[i]], result[vowelIndices[j]] = result[vowelIndices[j]], result[vowelIndices[i]]
        i++
        j--
    }

    return string(result)
}

func isVowel(rn rune) bool {
    return rn == 'a' || rn == 'e' || rn == 'i' || rn == 'o' || rn == 'u' || rn == 'A' || rn == 'E' || rn == 'I' || rn == 'O' || rn == 'U'
}