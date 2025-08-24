func lengthOfLIS(nums []int) int {
    tops := []int{ nums[0] }

    for _, num := range nums[1:] {
        if num > tops[len(tops) - 1] {
            tops = append(tops, num)
            continue
        }

        for j, top := range tops {
            if top >= num {
                tops[j] = num
                break
            }
        }
    }

    return len(tops)
}
