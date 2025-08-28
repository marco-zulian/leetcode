func minMeetingRooms(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
    })

    rooms := []int{}
    
    for _, interval := range intervals {
        foundRoom := false

        for i, occupiedUntil := range rooms {
            if occupiedUntil <= interval[0] {
                rooms[i] = interval[1]
                foundRoom = true
                break
            }
        }
        
        if !foundRoom {
            rooms = append(rooms, interval[1])
        }
    }
    
    return len(rooms)
}