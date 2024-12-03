func checkIfExist(arr []int) bool {
    seen := make(map[int]bool)

    for _, el := range arr {
        half := false
        if el % 2 == 0 {
            _, half = seen[el / 2]
        }
        _, double := seen[el * 2]
        
        if half || double {
            return true
        }

        seen[el] = true
    }

    return false
}