/*
 0, 1, 2, 3, 4, 5, 6
 0, 6, 1, 5, 2, 4, 3
 
 0, 1, 2, 3, 4, 5
 0, 5, 1, 4, 2, 3
 */
func reorderList(head *ListNode)  {
    first := head
    reorderListHelper(head, &first)
}

func reorderListHelper(second *ListNode, first **ListNode) bool {
    if second == nil {
        return false
    }
    if reorderListHelper(second.Next, first) {
        return true
    }
    if *first == second || (*first).Next == second {
        second.Next = nil
        return true
    }
    nextFirst := (*first).Next
    (*first).Next = second
    second.Next = nextFirst
    *first = nextFirst
    return false
}
