func findMin(nums []int) int {
    i, j := 0, len(nums) - 1

    for i <= j {
        k := (i + j) / 2

        if k == 0 {
            if len(nums) == 1 || nums[0] < nums[1] {
                return nums[0]
            } else {
                return nums[1]
            }
        }
        if nums[k] <= nums[0] && nums[k-1] >= nums[0] {
            return nums[k]
        } else if nums[k] <= nums[0] && nums[k-1] <= nums[k-1] {
            j = k - 1
        } else {
            i = k + 1
        }
    }

    return nums[0]
}