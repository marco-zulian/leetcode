func longestConsecutive(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }

    slices.Sort(nums)

    max := 0
    curr := 1
    for i, num := range nums[1:] {
        if num == nums[i] + 1 {
            curr++
        } else if num == nums[i] { 
            continue
        } else {
            if curr > max {
                max = curr
            }
            curr = 1
        }
    }

    if curr > max {
        max = curr
    }
    return max
}