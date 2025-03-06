func checkXMatrix(grid [][]int) bool {
    for i, row := range grid {
        for j, cell := range row {
            isOnDiagonal := isOnDiagonal(i, j, len(grid));
            if isOnDiagonal && cell == 0 {
                return false
            } else if !isOnDiagonal && cell != 0 {
                return false
            }
        }
    }

    return true
}

func isOnDiagonal(i int, j int, size int) bool {
    if i == j { return true } // main diagonal
    if i + j == size - 1 { return true } // opposite diagonal
    return false
}