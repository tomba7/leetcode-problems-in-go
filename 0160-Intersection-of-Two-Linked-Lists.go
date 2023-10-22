/**
    m [5,6,1,8,4,5]
    n   [4,1,8,4,5]
    m-n=x
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    m, n := listLen(headA), listLen(headB)
    if m < n { return getIntersectionNode(headB, headA) }
    diff := m - n
    for i := 0; i < diff; i++ {
        headA = headA.Next
    }
    for headA != nil && headB != nil {
        if headA == headB {
            return headA
        }
        headA, headB = headA.Next, headB.Next
    }
    return nil
}

func listLen(l *ListNode) int {
    var length int
    for curr := l; curr != nil; curr = curr.Next {
        length++
    }
    return length
}
