func maxScore(s string) int {
    score := 0

    if s[0] == '0' { score += 1 }

    for _, rn := range s[1:] {
        if rn == '1' {
            score += 1
        } 
    }

    maxScore := score

    for _, rn := range s[1:len(s)-1] {
        if rn == '0' {
            score += 1
            if score > maxScore {
                maxScore = score
            }
        } else {
            score -= 1
        }
    } 

    return maxScore
}