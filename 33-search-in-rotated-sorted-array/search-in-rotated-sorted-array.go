func search(nums []int, target int) int {
    var cut int
    i, j := 0, len(nums) - 1

    for i <= j {
        k := (i + j) / 2

        if k == 0 {
            if len(nums) == 1 || nums[0] < nums[1] {
                cut = 0
            } else {
                cut = 1
            }
            break
        }

        if nums[k-1] > nums[k] {
            cut = k
            break
        } else if nums[k] > nums[0] {
            i = k+1
        } else {
            j = k-1
        }
    }

    if target == nums[0] {
        return 0
    } else if target < nums[0] || cut == 0 {
        i, j = cut, len(nums) - 1
    } else {
        i, j = 0, cut-1
    }

    for i <= j {
        k := (i + j) / 2

        if nums[k] == target {
            return k
        } else if nums[k] < target {
            i = k + 1
        } else {
            j = k - 1
        }
    }

    return -1
}