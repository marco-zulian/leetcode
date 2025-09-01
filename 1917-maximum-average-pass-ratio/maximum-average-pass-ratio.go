import "container/heap"

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
    classesWithVal := make([]Values, len(classes))
    for i, v := range classes {
        classesWithVal[i] = Values{
            Curr: v,
            Gain: increaseWithExtraStudent(v),
        }
    }

    h := &MaxHeap{}
	*h = classesWithVal
	heap.Init(h)

    for _ = range extraStudents {
        val := heap.Pop(h).(Values)
        val.Curr[0]++
        val.Curr[1]++
        val.Gain = increaseWithExtraStudent(val.Curr)
        heap.Push(h, val)
    }

    sum := 0.0
    for _, v := range classesWithVal {
        sum += float64(v.Curr[0]) / float64(v.Curr[1])
    }
    return sum / float64(len(classes))
}

func increaseWithExtraStudent(curr []int) float64 {
    currPerc := float64(curr[0]) / float64(curr[1])
    plusOnePerc := (float64(curr[0]) + 1) / (float64(curr[1]) + 1)
    
    return plusOnePerc - currPerc
}

type Values struct {
    Curr []int
    Gain float64
}

type MaxHeap []Values
func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(Values))
}

func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func (h MaxHeap) Len() int {
    return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
    return h[i].Gain > h[j].Gain
}

func (h MaxHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}
