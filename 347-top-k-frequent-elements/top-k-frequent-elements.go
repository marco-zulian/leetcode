func topKFrequent(nums []int, k int) []int {
    counts := map[int]int{}
    frequencies := map[int][]int{}
    result := []int{}
    remaining := k

    for _, num := range nums {
        counts[num]++
    }

    for num, count := range counts {
        frequencies[count] = append(frequencies[count], num)
    }

    for i := range len(nums) {
        if nums, ok := frequencies[len(nums)-i]; ok {
            result = append(result, nums...)
            remaining -= len(nums)
            if remaining <= 0 { break }
        }
    }

    return result
}