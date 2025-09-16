func hasSameDigits(s string) bool {
    vals := []int{}

    for _, rn := range s {
        val, _ := strconv.Atoi(string(rn))
        vals = append(vals, val)
    }

    newVals := []int{}
    for len(vals) > 2 {
        for i := range len(vals) - 1 {
            newVals = append(newVals, (vals[i] + vals[i+1]) % 10)
        }

        vals = newVals
        newVals = []int{}
    }

    return vals[0] == vals[1]
}