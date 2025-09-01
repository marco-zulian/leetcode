func numIslands(grid [][]byte) int {
    islands := 0
    var helper func(int, int)
    
    helper = func(i, j int) {
        if i >= len(grid) || j >= len(grid[0]) || j < 0 || i < 0 { return }
        if grid[i][j] == byte('1') {
            grid[i][j] = byte('.')
            helper(i-1, j)
            helper(i+1, j)
            helper(i, j+1)
            helper(i, j-1)
        }
    }
    
    for i, row := range grid {
        for j, cell := range row {
            if cell == byte('1') {
                islands += 1
                grid[i][j] = byte('.')
                helper(i-1, j)
                helper(i+1, j)
                helper(i, j+1)
                helper(i, j-1)
            }
        }
    }
    
    return islands
}