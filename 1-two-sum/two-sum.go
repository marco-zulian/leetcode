func twoSum(nums []int, target int) []int {
    missingValues := map[int]int{}

    for i, num := range nums {
        if j, ok := missingValues[num]; ok {
            return []int{i, j}
        }
        missingValues[target - num] = i
    }

    return []int{-1, -1}
}