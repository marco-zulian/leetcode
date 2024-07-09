func findMatrix(nums []int) [][]int {
    var result [][]int
    acc := make(map[int]int)

    for _, v := range nums {
        if acc[v] == len(result) {
            result = append(result, []int{})
        }

        result[acc[v]] = append(result[acc[v]], v)
        acc[v]++
    }

    return result
}