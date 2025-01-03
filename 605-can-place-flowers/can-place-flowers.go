func canPlaceFlowers(flowerbed []int, n int) bool {
    added := 0
    if len(flowerbed) == 1 {
        if flowerbed[0] == 0 {
            added += 1
        }
        return added >= n
    }

    for i := range len(flowerbed) {
        if flowerbed[i] == 1 { continue }

        if i == 0 {
            if flowerbed[i+1] == 0 {
                flowerbed[i] = 1
                added += 1
            }
        } else if i == len(flowerbed) - 1 {
            if flowerbed[i-1] == 0 {
                flowerbed[i] = 1
                added += 1 
            }
        } else if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
            flowerbed[i] = 1
            added += 1
        }

        if added >= n {
            return true
        }
    }

    return added >= n
}