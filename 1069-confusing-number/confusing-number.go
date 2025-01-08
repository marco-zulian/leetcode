func confusingNumber(n int) bool {
    originalValue, flippedValue := n, 0

    for n > 0 {
        digit := n % 10
        
        if isValid(digit) {
            n = n / 10
            flippedValue = flippedValue * 10 + opositeValue(digit)
            continue
        }

        return false
    }

    return flippedValue != originalValue
}

func opositeValue(digit int) int {
    if digit == 0 { return 0 }
    if digit == 1 { return 1 }
    if digit == 6 { return 9 }
    if digit == 8 { return 8 }
    if digit == 9 { return 6 }
    
    return -1
}

func isValid(digit int) bool {
    return digit == 0 || digit == 1 || digit == 6 || digit == 8 || digit == 9
}