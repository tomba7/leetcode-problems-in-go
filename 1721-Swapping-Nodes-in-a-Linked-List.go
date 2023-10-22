func swapNodes(head *ListNode, k int) *ListNode {
    var listSize int
    var front, end *ListNode
    for curr := head; curr != nil; curr = curr.Next {
        listSize++
        if end != nil {
            end = end.Next
        }
        if listSize == k {
            front = curr
            end = head
        }
    }
    front.Val, end.Val = end.Val, front.Val
    return head
}

---

// 2 pass
func swapNodes(head *ListNode, k int) *ListNode {
    var listSize int
    var front *ListNode
    for curr := head; curr != nil; curr = curr.Next {
        listSize++
        if listSize == k {
            front = curr
        }
    }
    end := head
    for i := 1; i <= listSize-k; i++ {
        end = end.Next
    }
    front.Val, end.Val = end.Val, front.Val
    return head
}

---

// 3 pass
func swapNodes(head *ListNode, k int) *ListNode {
    var listSize int
    for curr := head; curr != nil; curr = curr.Next {
        listSize++
    }
    front := head
    for i := 1; i < k; i++ {
        front = front.Next
    }
    end := head
    for i := 1; i <= listSize-k; i++ {
        end = end.Next
    }
    front.Val, end.Val = end.Val, front.Val
    return head
}
