func minSwaps(data []int) int {
    totalOnes := 0

    for _, n := range data {
        if n == 1 {
            totalOnes++
        }
    }

    zerosIn, minSwaps := 0, len(data)
    for i, n := range data {
        if n == 0 {
            zerosIn++
        }
    
        if i + 1 >= totalOnes {
            if i - totalOnes >= 0 && data[i - totalOnes] == 0 {
                zerosIn--
            }

            if zerosIn < minSwaps {
                minSwaps = zerosIn
            }
        }
    }

    return minSwaps
}