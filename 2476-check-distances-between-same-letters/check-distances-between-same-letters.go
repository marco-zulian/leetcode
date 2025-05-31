func checkDistances(s string, distance []int) bool {
    positions := map[rune]int{}

    for i, rn := range s {
        if prevPos, ok := positions[rn]; ok {
            if distance[byte(rn) - byte('a')] != (i - prevPos - 1) {
                return false
            }
        }

        positions[rn] = i
    }

    return true
}