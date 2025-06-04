func coloredCells(n int) int64 {
    if n == 1 {
        return 1
    }

    if n == 2 {
        return 5
    }

    cellsPaintedOnLastExpansion := 4 * (n - 1)
    cumulativeCellsPainted := int64((5 + cellsPaintedOnLastExpansion) * (n - 1) / 2)
    return cumulativeCellsPainted - int64((n - 3) / 2)
}