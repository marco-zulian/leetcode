func averageValue(nums []int) int {
    var sum, count int

    for _, num := range nums {
        if num % 2 == 0 && num % 3 == 0 {
            sum += num
            count += 1
        }
    }

    if (count == 0) { return 0 }
    return sum / count
}