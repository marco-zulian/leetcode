func maximumScore(nums []int, multipliers []int) int {
    memo := map[string]int{}
    var maxHelper func(int, int, int) int
    
    maxHelper = func(op int, start int, end int) int {
        if len(multipliers) == op { return 0 }
        key := strconv.Itoa(op) + "==" + strconv.Itoa(start) + "==" + strconv.Itoa(end)
        if v, ok := memo[key]; ok { return v }
        
        memo[key] = max(
            multipliers[op] * nums[start] + maxHelper(op+1, start+1, end),
            multipliers[op] * nums[end] + maxHelper(op+1, start, end-1),
        )

        return memo[key] 
    }
    
    return maxHelper(0, 0, len(nums)-1)
}

func max(a int, b int) int {
    if a > b { return a }
    return b
}