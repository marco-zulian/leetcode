func minimumAbsDifference(arr []int) [][]int {
    result := [][]int{}
    min, max := 1000001, -1000001

    for _, n := range arr {
        if n > max { max = n }
        if n < min { min = n }
    }

    buckets := make([]int, max - min + 1)
    for _, n := range arr {
        buckets[n-min]++
    }

    prev := -50000000
    minDiff := 2000001
    for i, v := range buckets {
        if v == 0 { continue }
        distance := i + min - prev

        if distance < minDiff {
            minDiff = distance
            result = [][]int{[]int{prev, min+i}}
        } else if distance == minDiff {
            result = append(result, []int{prev, min+i})
        }

        prev = i + min
    }

    return result
}
