func longestCommonPrefix(strs []string) string {
    minLen := 201
    
    for _, str := range strs {
        if len(str) < minLen {
            minLen = len(str)
        }
    }
    
    for i := range minLen {
        curr := strs[0][i]
        for _, str := range strs[1:] {
            if str[i] != curr {
                return strs[0][:i]
            }
        }
    }
    
    return strs[0][:minLen]
}