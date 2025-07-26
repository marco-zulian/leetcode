func longestAlternatingSubarray(nums []int, threshold int) int {
    l, r, longest := 0, 0, 0

    for i, num := range nums {
        if num > threshold { 
            l, r = i+1, i+1
            continue
        }

        if i == l {
            if num % 2 == 0 {
                r++
                if r - l > longest {
                    longest = r - l
                }
            } else {
                l, r = i+1, i+1
            }
        }

        if num % 2 == (r - l) % 2 {
            r++
            if r - l > longest {
                longest = r - l
            }
        } else {
            if num % 2 == 0 {
                l, r = i, i+1
            } else {
                l, r = i+1, i+1
            }
        }
    }

    return longest
}