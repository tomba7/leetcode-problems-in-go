func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
    graph := newGraph(n, flights)
    distances := make([]int, n)
    for i := 1; i < len(distances); i++ {
        distances[i] = math.MaxInt32
    }
    currentStops := make([]int, n)
    for i := 1; i < len(currentStops); i++ {
        currentStops[i] = math.MaxInt32
    }
    h := &MinHeap{}
    heap.Push(h, []int{src, 0, 0})
    for h.Len() != 0 {
        info := heap.Pop(h).([]int)
        curr, cost, stops := info[0], info[1], info[2]
        if curr == dst { return cost }
        if stops == K + 1 { continue }
        for neighbor := 0; neighbor < n; neighbor++ {
            if graph[curr][neighbor] > 0 {
                dU, dV := cost, distances[neighbor]
                wUV := graph[curr][neighbor]
                if dU + wUV < dV {
                    heap.Push(h, []int{neighbor, dU + wUV, stops + 1})
                    distances[neighbor] = dU + wUV
                } else if stops < currentStops[neighbor] {
                    heap.Push(h, []int{neighbor, dU + wUV, stops + 1})
                    currentStops[neighbor] = stops
                }
            }
        }
    }
    if distances[dst] == math.MaxInt32 { return -1 }
    return distances[dst]
}

func newGraph(n int, flights [][]int) [][]int {
    graph := make([][]int, n)
    for i := range graph {
        graph[i] = make([]int, n)
    }
    for _, flight := range flights {
        graph[flight[0]][flight[1]] = flight[2]
    }
    return graph
}

type MinHeap [][]int
func (h MinHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Len() int { return len(h) }
func (h *MinHeap) Push(val interface{}) { *h = append(*h, val.([]int)) }
func (h *MinHeap) Pop() interface{} { x := (*h)[h.Len() - 1]; *h = (*h)[:h.Len() - 1]; return x }
