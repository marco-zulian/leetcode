func maximumScore(nums []int, multipliers []int) int {
    memo := map[string]int{}
    var maxHelper func([]int, []int, int, int) int
    
    maxHelper = func(nums []int, multipliers []int, start int, end int) int {
        if len(multipliers) == 0 { return 0 }
        key := strconv.Itoa(len(multipliers)) + "==" + strconv.Itoa(start) + "==" + strconv.Itoa(end)
        if v, ok := memo[key]; ok { return v }
        
        memo[key] = max(
            multipliers[0] * nums[0] + maxHelper(nums[1:], multipliers[1:], start+1, end),
            multipliers[0] * nums[len(nums) - 1] + maxHelper(nums[:len(nums)-1], multipliers[1:], start, end+1),
        )

        return memo[key] 
    }
    
    return maxHelper(nums, multipliers, 0, 0)
}

func max(a int, b int) int {
    if a > b { return a }
    return b
}