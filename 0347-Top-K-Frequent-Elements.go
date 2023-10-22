// Given a non-empty array of integers, return the k most frequent elements.
//
// For example,
// Given [1,1,1,2,2,3] and k = 2, return [1,2].
//
// Note:
// You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
// Your algorithm's time complexity must be better than O(n log n), where n is the array's size.

func topKFrequent(nums []int, k int) []int {
    freqMap, h, result, i := make(map[int]int), &MinHeap{}, []int{}, 0
    for _, n := range nums { freqMap[n]++ }
    if k <= 0 || k > len(freqMap) { return nil }
    for n, f := range freqMap {
        if i < k {
            heap.Push(h, &Entry{n, f})
        } else {
            if f < (*h)[0].freq { continue }
            heap.Pop(h)
            heap.Push(h, &Entry{n, f})
        }
        i++
    }
    for h.Len() != 0 {
        result = append(result, heap.Pop(h).(*Entry).value)
    }
    return result
}

type Entry struct { value, freq int }
type MinHeap []*Entry
func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(value interface{}) { *h = append(*h, value.(*Entry)) }
func (h *MinHeap) Pop() interface{} { x := (*h)[len(*h) - 1]; *h = (*h)[:len(*h) - 1]; return x }
