func summaryRanges(nums []int) []string {
    if len(nums) == 0 { return []string{} }

    startNum := nums[0]
    currNum := nums[0]
    result := []string{}

    for _, num := range nums[1:] {
        if num - currNum > 1 {
            result = appendRange(result, startNum, currNum)
            startNum, currNum = num, num
        } else {
            currNum = num
        }
    }

    result = appendRange(result, startNum, currNum)
    return result
}

func appendRange(result []string, startNum int, endNum int) []string {
    if startNum == endNum {
        result = append(result, fmt.Sprintf("%d", endNum))
    } else {
        result = append(result, fmt.Sprintf("%d->%d", startNum, endNum))
    }

    return result
}