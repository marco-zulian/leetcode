func longestConsecutive(nums []int) int {
    if len(nums) <= 1 { return len(nums) }
    
    set := map[int]bool{}

    for _, num := range nums {
        set[num] = true
    }

    max := 0
    curr := 1
    for num := range set {
        if _, ok := set[num-1]; ok { continue }

        val := num
        for {
            if _, ok := set[val + 1]; ok {
                curr++
                val++
            } else {
                if curr > max {
                    max = curr
                }
                curr = 1
                break
            }
        }

    }

    if curr > max {
        max = curr
    }
    return max
}