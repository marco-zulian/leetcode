func climbStairs(n int) int {
    if n == 1 { return 1 }
    if n == 2 { return 2 }

    result := make([]int, n)
    result[0] = 1
    result[1] = 2

    for i := range n-2 {
        result[i+2] = result[i] + result[i+1]
    }

    return result[n-1]
}