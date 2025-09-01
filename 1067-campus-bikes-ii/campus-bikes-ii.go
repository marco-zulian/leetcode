func assignBikes(workers [][]int, bikes [][]int) int {
    var helper func(int) int
    memo :=  map[string]int{}
    bikesTaken := map[int]bool{}

    getHash := func() string {
        hash := ""
        for i := range bikes {
            if v, _ := bikesTaken[i]; !v { hash += strconv.Itoa(i) }
        }
        return hash
    }

    helper = func(worker int) int {
        if worker == len(workers) - 1 {
            minDist := 5000
            for i, bike := range bikes {
                if v, _ := bikesTaken[i]; v { continue }
                dist := manhatamDistance(bike, workers[worker])
                if dist < minDist {
                    minDist = dist
                }
            }

            return minDist
        }
        
        if v, ok := memo[getHash()]; ok {
            return v
        }

        minVal := 1000000000
        for i, bike := range bikes {
            if v, _ := bikesTaken[i]; v { continue }
            bikesTaken[i] = true
            totalDist := manhatamDistance(bike, workers[worker]) + helper(worker+1)
            bikesTaken[i] = false

            if totalDist < minVal {
                minVal = totalDist
            }
        }

        memo[getHash()] = minVal
        return minVal
    }

    return helper(0)
}

func manhatamDistance(a []int, b []int) int {
    return Abs(a[0] - b[0]) + Abs(a[1] - b[1])
}

func Abs(a int) int {
    if a >= 0 { return a }
    return -a
}