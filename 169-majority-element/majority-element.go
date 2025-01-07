func majorityElement(nums []int) int {
    counts := map[int]int{}
    majority := int(math.Floor(float64(len(nums)) / 2))

    for _, num := range nums {
        counts[num]++
    }

    for k,v := range counts {
        if v > majority {
            return k
        }
    }

    return -1
}