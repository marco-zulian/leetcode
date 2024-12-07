func removeElement(nums []int, val int) int {
    last := len(nums) - 1
    var i int

    for i <= last {
        if nums[last] == val {
            last--
        } else if nums[i] == val {
            nums[i], nums[last] = nums[last], nums[i]
            last--
        } else {
            i++
        }
    }

    return last + 1
}