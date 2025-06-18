func maxArea(height []int) int {
    i, j := 0, len(height) - 1
    maxArea := 0

    for i < j {
        area := float64(j- i) * math.Min(float64(height[i]), float64(height[j]))
        maxArea = int(math.Max(float64(maxArea), area))

        if height[i] < height[j] {
            i++
        } else {
            j--
        }
    }

    return maxArea
}