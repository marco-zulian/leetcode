func countSubarrays(nums []int, k int) int64 {
    max := 0

    for _, num := range nums {
        if num > max {
            max = num
        }
    }


    i, subarrays, elements := 0, 0, 0

    for _, num := range nums {
        if num == max {
            elements += 1
        }

        if elements > k {
            for nums[i] != max {
                i += 1
            }

            elements -= 1
            i += 1
        }
    
        if elements == k {
            for nums[i] != max {
                i += 1
            }
    
            subarrays += (i + 1)
        }
    }

    return int64(subarrays)
}