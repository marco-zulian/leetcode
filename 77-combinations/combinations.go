func combine(n int, k int) [][]int {
    result := [][]int{}
    curr := []int{}
    
    var combineHelper func(int, int, int)
    combineHelper = func(n int, k int, i int) {
        if i > n { return }

        curr = append(curr, i)

        if len(curr) == k {
            result = append(result, append([]int{}, curr...))
        } else {
            for j := range n - i {
                combineHelper(n, k, i+j+1)
            }
        }
        
        curr = curr[:len(curr)-1]
        return
    }
    
    for i := range n {
        combineHelper(n,k,i+1)
    }
    
    return result
}