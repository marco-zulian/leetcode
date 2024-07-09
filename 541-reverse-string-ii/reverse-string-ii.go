func reverseStr(s string, k int) string {
    bytes := []byte(s)

    for i := 2 * k; i <= len(bytes); i += 2 * k {
        slices.Reverse(bytes[i - 2*k : i - k])
    }

    remainder := len(bytes) % (2*k)
    if (remainder < k) {
        slices.Reverse(bytes[len(bytes) - remainder:])
    } else if (remainder < 2*k) {
        slices.Reverse(bytes[len(bytes) - remainder: len(bytes) - remainder + k])
    }

    return string(bytes)
}
