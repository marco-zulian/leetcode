func dominantIndex(nums []int) int {
    largest, secondLargest, largestIndex := nums[0], -1, 0
    
    for i, num := range nums[1:] {
        if num > largest {
            secondLargest = largest
            largest = num
            largestIndex = i + 1
            continue
        }
        
        if num > secondLargest {
            secondLargest = num
        }
    }
    
    if largest < 2 * secondLargest { return - 1}
    return largestIndex
}