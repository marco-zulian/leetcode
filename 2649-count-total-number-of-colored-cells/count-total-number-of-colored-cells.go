func coloredCells(n int) int64 {
    return 1 + 4 * int64(n) * int64(n-1) / 2
}