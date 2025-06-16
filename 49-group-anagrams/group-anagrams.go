func groupAnagrams(strs []string) [][]string {
    anagrams := map[string][]string{}

    for _, str := range strs {
        wordLetters := map[rune]int{}

        for _, rn := range str {
            wordLetters[rn]++
        }

        letterWord := getLetterString(wordLetters)
        anagrams[letterWord] = append(anagrams[letterWord], str)
    }

    result := [][]string{}
    for _, value := range anagrams {
        result = append(result, value)
    }

    return result
}

func getLetterString(letterCount map[rune]int) string {
    word := ""

    keys := make([]rune, len(letterCount))
    for key := range letterCount {
        keys = append(keys, key)
    }

    slices.Sort(keys)
    for _, key := range keys {
        word += strconv.Itoa(letterCount[key]) + string(key)
    }

    return word
}