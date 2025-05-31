func countNegatives(grid [][]int) int {
    total := 0
    min := len(grid[0])

    for i := range grid {
        for j, v := range grid[i] {
            if j >= min {
                total += len(grid[0]) - min
                break
            }
            if v < 0 {
                min = j
                total += len(grid[0]) - min
                break
            }
        }
    }

    return total
}