func letterCombinations(digits string) []string {
    mapping := map[rune][]string{
        '2': []string{"a", "b", "c"},
        '3': []string{"d", "e", "f"},
        '4': []string{"g", "h", "i"},
        '5': []string{"j", "k", "l"},
        '6': []string{"m", "n", "o"},
        '7': []string{"p", "q", "r", "s"},
        '8': []string{"t", "u", "v"},
        '9': []string{"w", "x", "y", "z"},
    }
    
    if len(digits) == 0 { return []string{} }
    if len(digits) == 1 {
        result, _ := mapping[rune(digits[0])]
        return result
    }
    
    result := []string{}
    queue, _ := mapping[rune(digits[0])]
    
    for len(queue) != 0 {
        curr := queue[0]
        queue = queue[1:]
        nextPossibleRunes, _ := mapping[rune(digits[len(curr)])]
        
        if len(curr) == len(digits) - 1 {
            for _, rn := range nextPossibleRunes {
                result = append(result, curr + rn)
            }
        } else {
            for _, rn := range nextPossibleRunes {
                queue = append(queue, curr + rn)
            }
        }
    }
    
    return result
}