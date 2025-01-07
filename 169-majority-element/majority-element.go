func majorityElement(nums []int) int {
    counts := map[int]int{}
    majority := len(nums) / 2

    for _, num := range nums {
        counts[num]++
        if counts[num] > majority {
            return num
        }
    }

    return -1
}