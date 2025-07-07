func minSwaps(nums []int) int {
    var total, maxGrouped, currGrouped int

    for _, v := range nums {
        if v == 1 {
            total++
        }
    }

    arr := nums[0:total]
    for _, v := range arr {
        if v == 1 {
            currGrouped++
        }
    }
    maxGrouped = currGrouped

    for i := range nums {
        if nums[i] == 1 {
            currGrouped--
        }

        if nums[(i+total) % len(nums)] == 1 {
            currGrouped++
        }
        
        if currGrouped > maxGrouped {
            maxGrouped = currGrouped
        }
    }

    return total - maxGrouped
}