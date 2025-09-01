func findMissingRanges(nums []int, lower int, upper int) [][]int {
    if len(nums) == 0 { return [][]int{[]int{lower, upper}}}
    result := [][]int{}
    
    for i, num := range nums {
        if i == 0 {
            if num != lower {
                result = append(result, []int{lower, num-1})
            }
            continue
        }
        
        if num == nums[i-1] + 1 { continue }
        result = append(result, []int{nums[i-1]+1, num-1})
    }
    
    if nums[len(nums)-1] != upper {
        result = append(result, []int{nums[len(nums)-1]+1, upper})
    }

    return result
}