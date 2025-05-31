func largestUniqueNumber(nums []int) int {
    once := map[int]bool{}

    for _, n := range nums {
        _, ok := once[n]
        once[n] = !ok
    }

    mx := -1
    for k, v := range once {
        if v {
            mx = max(mx, k)
        }
    }

    return mx
}