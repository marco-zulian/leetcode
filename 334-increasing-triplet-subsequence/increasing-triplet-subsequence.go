func increasingTriplet(nums []int) bool {
    a, b := math.Inf(1), math.Inf(1)

    for _, num := range nums {
        if float64(num) < a {
            a = float64(num)
        } else if float64(num) > a && float64(num) < b {
            b = float64(num)
        } else if float64(num) > a && float64(num) > b {
            return true
        }
    }

    return false
}