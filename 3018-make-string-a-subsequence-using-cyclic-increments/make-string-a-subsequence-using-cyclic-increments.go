func canMakeSubsequence(str1 string, str2 string) bool {
    bytes1 := []byte(str1)
    bytes2 := []byte(str2)

    var start int

    fmt.Println(bytes1)
    fmt.Println(bytes2)

    for _, v := range bytes2 {
        ok := false
        for start < len(bytes1) {
            if bytes1[start] == v || math.Max(97, float64((bytes1[start] + 1) % 123)) == float64(v) {
                start++
                ok = true
                break
            }
            start++
        }

        if !ok {
            return false
        }
    }

    return true
}