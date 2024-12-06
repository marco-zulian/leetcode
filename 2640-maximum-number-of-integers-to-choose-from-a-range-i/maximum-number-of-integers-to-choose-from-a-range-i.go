func maxCount(banned []int, n int, maxSum int) int {
    sum := 0
    values := 0
    banMap := make(map[int]bool)

    for _, ban := range banned {
        banMap[ban] = true
    }

    for val := range n + 1 {
        if val == 0 {
            continue
        }

        if sum + val > maxSum {
            break
        }

        if !banMap[val] {
            sum += val
            values += 1
        }
    }

    return values
}