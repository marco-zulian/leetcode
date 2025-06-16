func productExceptSelf(nums []int) []int {
    prefixProduct := 1
    sufixProducts := make([]int, len(nums))
    result := make([]int, len(nums))

    sufixProduct := nums[len(nums) - 1]
    for i := range(len(nums) - 1) {
        sufixProducts[len(nums) - 1 - i] = sufixProduct
        sufixProduct *= nums[len(nums) - 2 - i]
    }
    sufixProducts[0] = sufixProduct

    for i, num := range(nums) {
        if i == len(nums) - 1 {
            result[i] = prefixProduct
        } else {
            result[i] = prefixProduct * sufixProducts[i + 1]
            prefixProduct *= num
        }
    }

    return result
}