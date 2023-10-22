func mergeKLists(lists []*ListNode) *ListNode {
    var dummy ListNode
    prev := &dummy
    h := &MinHeap{}
    for i := 0; i < len(lists); i++ {
        if lists[i] == nil { continue }
        heap.Push(h, lists[i])
    }
    for h.Len() != 0 {
        node := heap.Pop(h).(*ListNode)
        prev.Next = node
        prev = prev.Next
        if node.Next == nil { continue }
        heap.Push(h, node.Next)
    }
    return dummy.Next
}

type MinHeap []*ListNode
func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Len() int { return len(h) }
func (h *MinHeap) Push(val interface{}) { *h = append(*h, val.(*ListNode)) }
func (h *MinHeap) Pop() interface{} { x := (*h)[h.Len() - 1]; *h = (*h)[:h.Len() - 1]; return x }

---

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 Start: 10:20 - 10:50
 L1:0 [1,4,5]
         ^
 L2:1 [1,3,4]
         ^
 L3:2 [2,6]
         ^
 Heap: (1, 1, 2)
 
 - Push the first elements of each list into a min heap
 - Until the heap is empty continue to pop and 
   add the value to the result
   followed by pushing the next element from the corresponding list
 */
func mergeKLists(lists []*ListNode) *ListNode {
    h := &MinHeap{}
    for i := 0; i < len(lists); i++ {
        if lists[i] == nil { continue }
        heap.Push(h, lists[i])
    }
    var dummy ListNode
    prev := &dummy
    for h.Len() != 0 {
        curr := heap.Pop(h).(*ListNode)
        prev.Next = &ListNode{Val: curr.Val}
        prev = prev.Next
        if curr.Next != nil {
            heap.Push(h, curr.Next)
        }
    }
    return dummy.Next
}

type MinHeap []*ListNode
func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }

func (h *MinHeap) Pop() interface{} {
    n := len(*h)
    x := (*h)[n-1]
    *h = (*h)[:n-1]
    return x
}
