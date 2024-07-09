func findPrefixScore(nums []int) []int64 {
    result := make([]int64, len(nums) + 1)
    var max int

    for i, v := range nums {
        if v > max {
            max = v
        }

        result[i + 1] = result[i] + int64(nums[i] + max)
    }

    return result[1:]
}