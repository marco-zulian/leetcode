func canJump(nums []int) bool {
    memo := map[int]bool{}
    
    var canReachEnd func(int, []int) bool
    canReachEnd = func(from int, nums []int) bool {
        if from == len(nums) - 1 { return true }
        if from >= len(nums) { return false }
        if canReach, ok := memo[from]; ok { return canReach }
        
        for i := 1; i <= nums[from]; i++ {
            if canReachEnd(from + i, nums) { return true }
        }

        memo[from] = false
        return false
    }
    
    return canReachEnd(0, nums)
}