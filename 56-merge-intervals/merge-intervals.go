func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    i := 0
    for i < len(intervals) - 1 {
        if intervals[i+1][0] > intervals[i][1] {
            i++
            continue
        }

        var newInterval []int
        if intervals[i+1][1] > intervals[i][1] {
            newInterval = []int{intervals[i][0], intervals[i+1][1]}    
        } else {
            newInterval = intervals[i]
        }
        
        intervals[i] = newInterval
        intervals = append(intervals[:i+1], intervals[i+2:]...)        
    }

    return intervals
}