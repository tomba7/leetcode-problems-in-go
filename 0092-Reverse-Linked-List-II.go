func reverseBetween(head *ListNode, left int, right int) *ListNode {
    var prev *ListNode
    curr := head
    for ; left > 1; left-- {
        prev = curr
        curr = curr.Next
        right--
    }
    first, second := prev, curr
    prev = nil
    for ; right > 0; right-- {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    second.Next = curr
    if first == nil {
        head = prev
    } else {
        first.Next = prev
    }
    return head
}
