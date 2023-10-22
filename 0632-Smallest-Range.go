func smallestRange(nums [][]int) []int {
    pq := &priorityQ{}
    var start, end int
    hi, minRange := math.MinInt32, math.MaxInt32
    for i := range nums {
        heap.Push(pq, pair{
            row: i, indx: 0, value: nums[i][0],
        })
        hi = max(hi, nums[i][0])
    }
    for pq.Len() == len(nums) {
        curr := heap.Pop(pq).(pair)
        if hi - curr.value < minRange {
            minRange = hi - curr.value
            start = curr.value
            end = hi
        }
        curr.indx++
        if curr.indx < len(nums[curr.row]) {
            curr.value = nums[curr.row][curr.indx]
            heap.Push(pq, curr)
            hi = max(hi, curr.value)
        }
    }
    return []int{start, end}
}

type pair struct { row, indx, value int }
type priorityQ []pair
func (pq priorityQ) Len() int { return len(pq) }
func (pq priorityQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQ) Less(i, j int) bool { return pq[i].value < pq[j].value }
func (pq *priorityQ) Push(x interface{}) { *pq = append(*pq, x.(pair)) }
func (pq *priorityQ) Pop() interface{} {
    x := (*pq)[len(*pq)-1]
    *pq = (*pq)[:len(*pq)-1]
    return x
}
func max(x, y int) int { if x > y { return x }; return y }
