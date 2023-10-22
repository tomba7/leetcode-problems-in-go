import "container/heap"

func connectSticks(sticks []int) int {
    h := &Heap{}
    for _, stick := range sticks {
        heap.Push(h, stick)
    }
    var res int
    for h.Len() > 1 {
        first, second := heap.Pop(h).(int), heap.Pop(h).(int)
        cost := first + second
        heap.Push(h, cost)
        res += cost
    }
    return res
}

type Heap []int
func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *Heap) Pop() interface{} { x := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return x }
