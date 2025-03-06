func percentageLetter(s string, letter byte) int {
    count := 0

    for _, rn := range s {
        if byte(rn) == letter {
            count++
        }
    }

    return (count * 100) / len(s)
}