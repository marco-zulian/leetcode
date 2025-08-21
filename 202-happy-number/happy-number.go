func isHappy(n int) bool {
    findNext := func(val int) int {
        next := 0

        for val > 0 {
            next += (val % 10) * (val % 10)
            val /= 10
        }

        return next
    }

    fast, slow := findNext(n), n
    
    for fast != slow {
        if fast == 1 || slow == 1 { return true }

        fast = findNext(findNext(fast))
        slow = findNext(slow)
    }

    return fast == 1
}