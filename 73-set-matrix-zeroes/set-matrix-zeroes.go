func setZeroes(matrix [][]int)  {
    swapped := make([][]bool, len(matrix))
    for i := range len(matrix) {
        swapped[i] = make([]bool, len(matrix[0]))
    }
    
    for i, row := range matrix {
        for j, cell := range row {
            if cell == 0 && !swapped[i][j] {
                for k := range row {
                    if matrix[i][k] != 0 {
                        matrix[i][k] = 0
                        swapped[i][k] = true   
                    }
                }
                
                for k := range matrix {
                    if matrix[k][j] != 0 {
                        matrix[k][j] = 0
                        swapped[k][j] = true
                    }
                }
            }
        }
    }
}