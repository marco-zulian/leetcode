func wiggleSort(nums []int)  {
    result := make([]int, len(nums))
    for i,v := range nums {
        result[i] = v
    }
    slices.Sort(result)  

    i, j, k := 0, len(nums)-1, 0
    for k < len(nums) {
        nums[k] = result[i]
        i++
        k++
        if k < len(nums) { nums[k] = result[j] }
        j--
        k++
    }
}