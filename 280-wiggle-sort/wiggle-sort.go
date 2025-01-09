func wiggleSort(nums []int)  {
    for i,_ := range nums[1:] {
        if i % 2 == 0 {
            if nums[i] > nums[i+1] {
                nums[i], nums[i+1] = nums[i+1], nums[i]
            }
        } else {
            if nums[i] < nums[i+1] {
                nums[i], nums[i+1] = nums[i+1], nums[i]
            }
        }
    }
}