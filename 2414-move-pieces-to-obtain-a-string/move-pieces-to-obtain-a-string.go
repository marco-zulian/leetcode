func canChange(start string, target string) bool {
    if len(start) != len(target) { return false }
    if strings.Replace(start, "_", "", -1) != strings.Replace(target, "_", "", -1) { return false }

    var j int
    for i, r := range target {
        if r == '_' { continue }
        for start[j] == '_' {
            j++
        }

        if r == 'R' && j > i { return false }
        if r == 'L' && j < i { return false }
        
        j++
    }

    return true
}