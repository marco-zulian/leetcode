func minimumAbsDifference(arr []int) [][]int {
    result := [][]int{}
    sort.Ints(arr)
    
    minDiff := 2000001
    for i := range arr[:len(arr)-1] {
        if arr[i+1] - arr[i] < minDiff {
            minDiff = arr[i+1] - arr[i]
            result = [][]int{[]int{arr[i], arr[i+1]}}
        } else if arr[i+1] - arr[i] == minDiff {
            result = append(result, []int{arr[i], arr[i+1]})
        }
    }
    
    return result
}