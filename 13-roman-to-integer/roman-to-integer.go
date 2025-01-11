func romanToInt(s string) int {
    currVal, currSum, result := 0, 0, 0
    conversions := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

    for _, rn := range s {
        val := conversions[rn]
        if val == currVal {
            currSum += val
        } else {
            if val < currVal {
                result += currSum
            } else {
                result -= currSum
            }
            currVal = val
            currSum = val
        }
    }

    return result + currSum
}