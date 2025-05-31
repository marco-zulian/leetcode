func countNegatives(grid [][]int) int {
    total := 0

    for i := range grid {
        for _, v := range grid[i] {
            if v < 0 {
                total++
            }
        }
    }

    return total
}