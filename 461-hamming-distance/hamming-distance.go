func hammingDistance(x int, y int) int {
    numberWithDifferentDigits := x ^ y
    hammingDistance := 0
    
    for numberWithDifferentDigits > 0 {
        if numberWithDifferentDigits % 2 == 1 {
            hammingDistance += 1
        }
        
        numberWithDifferentDigits /= 2
    }
    
    return hammingDistance
}