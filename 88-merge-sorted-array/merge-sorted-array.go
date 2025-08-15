func merge(nums1 []int, m int, nums2 []int, n int)  {
    if len(nums2) == 0 { return }

    for i := 1; i <= m; i++ {
        nums1[len(nums1)-i] = nums1[m-i]
        nums1[m-i] = 0
    }

    curr_j, curr_i := 0, n
    for i := 0; i < len(nums1); i++ {
        if curr_i < 0 || curr_i >= len(nums1) || (curr_j < n && nums2[curr_j] < nums1[curr_i]) {
            nums1[i] = nums2[curr_j]
            curr_j++
        } else {
            nums1[i] = nums1[curr_i]
            curr_i++
        }
    }
}