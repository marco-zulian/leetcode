func stringShift(s string, shift [][]int) string {
    totalShift := 0

    for _, op := range shift {
        if op[0] == 0 {
            totalShift -= op[1]
        } else {
            totalShift += op[1]
        }
    }

    shiftMod := totalShift % len(s)

    if shiftMod > 0 {
        return s[len(s)-shiftMod:] + s[:len(s)-shiftMod]
    } else {
        return s[(shiftMod * -1):] + s[:(shiftMod * -1)]
    }
}