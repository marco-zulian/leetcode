func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
    result := 0
    memo := map[string]int{}
    hash := func(sum int, level int) string {
        sumStr := strconv.Itoa(sum)
        levelStr := strconv.Itoa(level)
        return sumStr + "=>" + levelStr
    }
    
    for _, num1 := range nums1 {
        val := 0
        if v, ok := memo[hash(num1, 1)]; ok {
            result += v
            continue
        }
        
        for _, num2 := range nums2 {
            val2 := 0

            if v, ok := memo[hash(num1 + num2, 2)]; ok {
                val += v
                continue
            }

            for _, num3 := range nums3 {
                val3 := 0
                if v, ok := memo[hash(num1 + num2 + num3, 3)]; ok {
                    val2 += v
                    continue
                }

                for _, num4 := range nums4 {
                    if num1 + num2 + num3 + num4 == 0 {
                        val3++
                    }
                }
    
                memo[hash(num1 + num2 + num3, 3)] = val3
                val2 += val3
            }
            
            memo[hash(num1 + num2, 2)] = val2
            val += val2
        }
        
        memo[hash(num1, 1)] = val
        result += val
    }
    
    return result
}
