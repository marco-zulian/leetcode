func deleteAndEarn(nums []int) int {
    numValues := map[int]int{}
    memo := map[int]int{}
    maxNum, maxPoints := 0, 0
    
    var getValueOf func(int) int
    getValueOf = func(num int) int {
        if num <= 0 { return 0 }
        if num == 1 { return numValues[1] }
        if v, ok := memo[num]; ok { return v }
        
        memo[num] = max(getValueOf(num-1), numValues[num] + getValueOf(num-2))
        return memo[num]
    }
    
    for _, num := range nums {
        numValues[num] += num
        maxNum = max(num, maxNum)
    }
    
    for i := range maxNum {
        if _, ok := numValues[maxNum - i]; !ok { continue }
        maxPoints = max(maxPoints, getValueOf(maxNum - i))
    }
        
    return maxPoints
}

func max(a int, b int) int {
    if a > b { return a }
    return b
}