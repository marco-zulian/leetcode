func smallestNumber(n int) int {
    curr := 1

    for curr < (n + 1) {
        curr *= 2
    }

    return curr - 1
}