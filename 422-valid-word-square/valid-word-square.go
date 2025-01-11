func validWordSquare(words []string) bool {
    maxSize := len(words)
    for _, word := range words {
        if len(word) > maxSize {
            maxSize = len(word)
        }
    }


    wordsMatrix := make([][]rune, maxSize)
    for i, _ := range wordsMatrix {
        wordsMatrix[i] = make([]rune, maxSize)
    }

    for i, word := range words {
        for j, rn := range word {
            wordsMatrix[i][j] = rn
        }
    }

    for i, row := range wordsMatrix {
        for j, cell := range row {
            if i == j { break }
            if cell != wordsMatrix[j][i] { return false }
        }
    }

    return true
}