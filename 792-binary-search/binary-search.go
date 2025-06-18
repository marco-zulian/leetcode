func search(nums []int, target int) int {
    i, j := 0, len(nums) - 1

    for i <= j {
        k := (i + j) / 2

        if nums[k] == target {
            return k
        } else if nums[k] > target {
            j = k - 1
        } else {
            i = k + 1
        }
    }

    return -1
}