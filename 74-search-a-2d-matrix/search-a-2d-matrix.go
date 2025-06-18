func searchMatrix(matrix [][]int, target int) bool {
    i, j := 0, len(matrix) * len(matrix[0]) - 1
    
    for i <= j {
        k := (i + j) / 2

        row, column := k / len(matrix[0]), k % len(matrix[0])
        if matrix[row][column] == target {
            return true
        } else if matrix[row][column] < target {
            i = k + 1
        } else {
            j = k - 1
        }
    }

    return false
}