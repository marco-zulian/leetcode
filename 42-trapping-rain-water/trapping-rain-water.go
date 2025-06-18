func trap(height []int) int {
    var i, j, trap, trapPotential int

    for i < len(height) && j < len(height) {
        if height[j] >= height[i] {
            trap += trapPotential
            i = j
            j++
            trapPotential = 0
        } else {
            trapPotential += height[i] - height[j]
            j++
        }
    }

    i, j, trapPotential = len(height) - 1, len(height) - 1, 0
    for i >= 0 && j >= 0 {
        if height[j] > height[i] {
            trap += trapPotential
            i = j
            j--
            trapPotential = 0
        } else {
            trapPotential += height[i] - height[j]
            j--
        }
    }

    return trap
}