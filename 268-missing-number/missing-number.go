func missingNumber(nums []int) int {
    i := 0

    for i < len(nums) {
        if nums[i] == i { 
            i++
            continue
        }

        if nums[i] == len(nums) {
            i++
            continue
        }

        curr := nums[i]
        nums[curr], nums[i] = nums[i], nums[curr]
    }

    for i, num := range nums {
        if num != i {
            return i
        }
    }
    return len(nums)
}