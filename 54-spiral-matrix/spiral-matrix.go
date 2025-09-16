func spiralOrder(matrix [][]int) []int {
    i_min, i_max := 0, len(matrix)
    j_min, j_max := 0, len(matrix[0])
    
    result := make([]int, len(matrix) * len(matrix[0]))
    k := 0
    
    for i_min <= i_max && j_min <= j_max {
        for j := j_min; j < j_max && k < len(result); j++ {
            result[k] = matrix[i_min][j]
            k++
        }
    
        for i := i_min + 1; i < i_max && k < len(result); i++ {
            result[k] = matrix[i][j_max-1]
            k++
        }
        
        for j := j_max - 2; j >= j_min && k < len(result); j-- {
            result[k] = matrix[i_max - 1][j]
            k++
        }

        for i := i_max - 2; i > i_min && k < len(result); i-- {
            result[k] = matrix[i][j_min]
            k++
        }

        i_min++
        j_min++
        i_max--
        j_max--
    }
    
    return result
}