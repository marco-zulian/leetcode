func toHexspeak(num string) string {
    intVal, _ :=  strconv.Atoi(num);

    hex := fmt.Sprintf("%X", intVal);
    hex = strings.Replace(strings.Replace(hex, "1", "I", -1), "0", "O", -1);

    valids := map[rune]bool{'A': true, 'B': true, 'C': true, 'D': true, 'E': true, 'F': true, 'I': true, 'O': true};

    for _, rn := range hex {
        if _, ok := valids[rn]; !ok {
            return "ERROR"
        }
    }

    return hex
}