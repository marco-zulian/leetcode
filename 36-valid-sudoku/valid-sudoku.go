func isValidSudoku(board [][]byte) bool {
    rows := map[int]map[byte]bool{}
    cols := map[int]map[byte]bool{}
    boxes := map[int]map[byte]bool{}

    for i := range(9) {
        rows[i] = map[byte]bool{}
        cols[i] = map[byte]bool{}
        boxes[i] = map[byte]bool{}
    }

    for i, row := range(board) {
        for j, cell := range(row) {
            if string(cell) == "." { continue }
            if _, ok := rows[i][cell]; ok { return false }
            if _, ok := cols[j][cell]; ok { return false }
            if _, ok := boxes[3 * (i / 3) + (j / 3)][cell]; ok { return false }

            rows[i][cell] = true
            cols[j][cell] = true
            boxes[3 * (i / 3) + (j / 3)][cell] = true
        }
    }

    return true
}