func fib(n int) int {
    if n <= 1 { return n }

    fib := make([]int, n)
    fib[0] = 1
    fib[1] = 1

    for i := range n - 2 {
        fib[i+2] = fib[i] + fib[i+1]
    }

    return fib[n-1]
}