import "strconv"

func reverse(x int) int {
    str := strconv.Itoa(x)
    negative := false
    value := 0
    power := 1
    
    for _, rn := range str {
        if rn == '-' {
            negative = true
            continue
        }

        val, _ := strconv.Atoi(string(rn)) 
        
        if negative && value > (2147483648 - val * power ) { return 0 }
        if value > (2147483647 - val * power ) { return 0 }
        
        value += val * power
        power *= 10
    }

    if negative {
        return -value
    }
    return value
}