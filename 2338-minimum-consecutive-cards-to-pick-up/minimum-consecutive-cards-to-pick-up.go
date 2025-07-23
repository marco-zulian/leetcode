func minimumCardPickup(cards []int) int {
    seen := map[int]int{}
    j, sliding := 0, false

    for _, card := range cards {
        if sliding {
            seen[cards[j]]--
            j++
        }

        if pair, ok := seen[card]; ok && pair == 1 {
            for cards[j] != card {
                seen[cards[j]]--
                j++
            }
            sliding = true
        }

        seen[card]++
    }

    if !sliding { return - 1 }
    return len(cards) - j
}