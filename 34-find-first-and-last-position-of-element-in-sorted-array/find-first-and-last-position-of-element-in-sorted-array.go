func searchRange(nums []int, target int) []int {
    if len(nums) == 0 { return []int{-1, -1} }
    
    i, j := 0, len(nums) - 1
    start, end := -1, -1

    for i < j {
        mid := (i + j) / 2

        if nums[mid] < target {
            i = mid + 1
        } else {
            j = mid
        }
    }

    if nums[i] != target { return []int{-1, -1} }
    start = i

    i, j = 0, len(nums) - 1
    for i < j {
        mid := (i + j) / 2

        if nums[mid] <= target {
            i = mid + 1
        } else {
            j = mid
        }
    }

    if nums[i] == target {
        end = i
    } else {
        end = i - 1
    }
    return []int{start, end}
}