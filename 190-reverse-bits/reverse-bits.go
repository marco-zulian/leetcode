func reverseBits(n int) int {
    val := 0
    
    for _ = range 31 {
        val += n % 2
        val *= 2
        n /= 2
    }
    
    return val
}