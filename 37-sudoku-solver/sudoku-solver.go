func solveSudoku(board [][]byte)  {
    var getPossibleValues func(int, int) []byte
    var getCellValue func(int, int) bool

    getCellValue = func(row int, col int) bool {
        if board[row][col] != byte('.') {
            if row == 8 && col == 8 { return true }
            return getCellValue(row + ((col + 1) / 9), (col + 1) % 9)
        }
        
        for _, val := range getPossibleValues(row, col) {
            board[row][col] = val
            if row == 8 && col == 8 { return true }
            isFinished := getCellValue(row + ((col + 1) / 9), (col + 1) % 9)
            if isFinished { return true }
        }
        
        board[row][col] = byte('.')
        return false
    }
    
    getPossibleValues = func(row int, col int) []byte {
        possibilities := map[byte]bool{
            byte('1'): true, byte('2'): true, byte('3'): true, 
            byte('4'): true, byte('5'): true, byte('6'): true,
            byte('7'): true, byte('8'): true, byte('9'): true,
        }
        
        for i := range 9 {
            if board[i][col] != '.' { possibilities[board[i][col]] = false }
            if board[row][i] != '.' { possibilities[board[row][i]] = false }
        }
        
        startRow, startCol := row / 3, col / 3
        for i := range 3 {
            for j := range 3 {
                if board[startRow * 3 + i][startCol * 3 + j] != '.' { 
                    possibilities[board[startRow * 3 + i][startCol * 3 + j]] = false
                }
            }
        }
        
        result := []byte{}
        for k, v := range possibilities {
            if v {
                result = append(result, byte(k))
            }
        }
        return result
    }
    
    getCellValue(0, 0)
}