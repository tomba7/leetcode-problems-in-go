func minMeetingRooms(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    h := &Heap{}
    heap.Push(h, intervals[0][1])
    for _, interval := range intervals[1:] {
        if interval[0] >= h.Top() {
            heap.Pop(h)
        }
        heap.Push(h, interval[1])
    }
    return h.Len()
}

type Heap []int
func (h Heap) Top() int { return h[0] }
func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *Heap) Pop() interface{} { x := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return x }

---

func minMeetingRooms(intervals [][]int) int {
    var endPoints []EndPoint
    for _, interval := range intervals {
        endPoints = append(endPoints, EndPoint{ts: interval[0], marker: start})
        endPoints = append(endPoints, EndPoint{ts: interval[1], marker: end})
    }
    sort.Slice(endPoints, func(i, j int) bool {
        if endPoints[i].ts == endPoints[j].ts {
            return endPoints[i].marker == end && endPoints[j].marker == start
        }
        return endPoints[i].ts < endPoints[j].ts
    })
    result, roomCount := 0, 0
    for _, endPoint := range endPoints {
        if endPoint.marker == start {
            roomCount++
            result = max(result, roomCount)
        } else {
            roomCount--
        }
    }
    return result
}

const start, end = 0, 1
type EndPoint struct { ts, marker int }
func max(x, y int) int { if x > y {return x}; return y }
