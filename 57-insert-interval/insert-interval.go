func insert(intervals [][]int, newInterval []int) [][]int {
    if len(intervals) == 0 { return [][]int{newInterval} }

    mergedIntervals := [][]int{}
    i := 0

    for i < len(intervals) && intervals[i][1] < newInterval[0] {
        mergedIntervals = append(mergedIntervals, intervals[i])
        i++
    }

    for i < len(intervals) && intervals[i][0] <= newInterval[1] {
        if intervals[i][0] < newInterval[0] {
            newInterval[0] = intervals[i][0]
        }

        if intervals[i][1] > newInterval[1] {
            newInterval[1] = intervals[i][1]
        }

        i++
    }

    return append(append(mergedIntervals, newInterval), intervals[i:]...)
}
