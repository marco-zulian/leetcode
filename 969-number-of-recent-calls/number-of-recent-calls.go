type RecentCounter struct {
    times []int
}


func Constructor() RecentCounter {
    return RecentCounter{}
}


func (this *RecentCounter) Ping(t int) int {
    removed := 0

    for _, v := range this.times {
        if v < t - 3000 {
            removed++
        } else {
            break
        }
    }

    this.times = append(this.times[removed:], t)
    return len(this.times)
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */