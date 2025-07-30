func findDuplicate(nums []int) int {
    i := 0

    for i < len(nums) {
        if nums[i] == i + 1 {
            i++
            continue
        }

        curr := nums[i]
        if nums[curr - 1] == curr {
            return curr
        }

        nums[curr - 1], nums[i] = nums[i], nums[curr - 1]
    }

    return -1
}