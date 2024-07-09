func singleNonDuplicate(nums []int) int {
    i, j := 0, len(nums)
    var mid int

    fmt.Println(i, j)
    for i < j {
        mid = (i + j) / 2
        if mid == 0 || mid == len(nums) - 1 {
            return nums[mid]
        }

        if nums[mid - 1] == nums[mid] {
            if mid % 2 == 0 {
                j = mid
            } else {
                i = mid + 1
            }
        } else if nums[mid + 1] == nums[mid] {
            if (mid % 2 == 0) {
                i = mid + 1
            } else {
                j = mid
            }
        } else {
            return nums[mid]
        }
    }

    return -1
}