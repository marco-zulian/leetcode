func totalFruit(fruits []int) int {
    if len(fruits) <= 2 {
        return len(fruits)
    }

    var maxFruits, currFruits, fruitB, left, right, start int
    fruitA := fruits[0]

    for i, fruit := range fruits {
        if fruit != fruitA {
            fruitB = fruit
            left, right, maxFruits, currFruits, start = i, i + 1, i + 1, i + 1, i + 1
            break
        }
    }

    for k, fruitType := range fruits[start:] {
        if fruitType == fruitA {
            currFruits++
            fruitA, fruitB = fruitB, fruitA
            left, right = start + k, start + k + 1
        } else if fruitType == fruitB {
            currFruits++
            right++
        } else {
            fruitA, fruitB = fruitB, fruitType
            currFruits = right - left + 1
            left, right = start + k, start + k + 1
        }

        if currFruits > maxFruits {
            maxFruits = currFruits
        }
    }

    return maxFruits
}