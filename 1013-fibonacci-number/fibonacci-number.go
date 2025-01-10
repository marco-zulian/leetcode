func fib(n int) int {
    if n <= 1 {
        return n
    }

    a, b := 1, 1
    for _ = range n - 2 {
        a,b = b,a+b
    }

    return b
}