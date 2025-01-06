func minOperations(boxes string) []int {
    ballsDistances := make([]int, len(boxes))

    for i, box := range boxes {
        if box == '1' {
            for j, _ := range ballsDistances {
                ballsDistances[j] += int(math.Abs(float64(i-j)))
            }
        }
    }

    return ballsDistances
}