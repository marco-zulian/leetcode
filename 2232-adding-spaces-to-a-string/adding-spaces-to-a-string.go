func addSpaces(s string, spaces []int) string {
    var builder strings.Builder
    var currIndex int

    for _, index := range spaces {
        builder.WriteString(s[currIndex:index])
        builder.WriteString(" ")
        currIndex = index
    }

    builder.WriteString(s[currIndex:])
    return builder.String()
}