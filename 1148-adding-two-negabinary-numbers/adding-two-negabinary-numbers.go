func addNegabinary(arr1 []int, arr2 []int) []int {
    maxSize := max(len(arr1), len(arr2))
    var solution = make([]int, maxSize + 2)
    var biggestArray []int

    if len(arr1) >= len(arr2) {
        biggestArray = arr1
    } else {
        biggestArray = arr2
    }

    for i := 1; i <= min(len(arr1), len(arr2)); i++ {
        sum := arr1[len(arr1) - i] + arr2[len(arr2) - i] + solution[len(solution) - i]
        positionValue := sum % 2
        remainder := sum / 2

        solution[len(solution) - i] = positionValue
        if remainder == 1 {
            solution[len(solution) - i - 1] += 1
            solution[len(solution) - i - 2] += 1
        } else if remainder == 2 {
            solution[len(solution) - i - 2] += 1
        }

        if (solution[len(solution) - i - 2] == 1 && solution[len(solution) - i - 1] == 2) {
            solution[len(solution) - i - 1] = 0
            solution[len(solution) - i - 2] = 0
        }
    }

    for i := min(len(arr1), len(arr2)) + 1; i <= len(biggestArray); i++ {
        sum := biggestArray[len(biggestArray) - i] + solution[len(solution) - i]
        positionValue := sum % 2
        remainder := sum / 2

        solution[len(solution) - i] = positionValue
        if remainder == 1 {
            solution[len(solution) - i - 1] += 1
            solution[len(solution) - i - 2] += 1
        } else if remainder == 2 {
            solution[len(solution) - i - 2] += 1
        }

        if (solution[len(solution) - i - 2] == 1 && solution[len(solution) - i - 1] == 2) {
            solution[len(solution) - i - 1] = 0
            solution[len(solution) - i - 2] = 0
        }
    }

    var startPosition int
    for startPosition < len(solution) - 1 && solution[startPosition] == 0 {
        startPosition++
    }

    return solution[startPosition:]
}
