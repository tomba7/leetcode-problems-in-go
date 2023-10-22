/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    length := getLength(head)
    result, _ := isPalindromeHelper(head, length)
    return result
}

func isPalindromeHelper(node *ListNode, length int) (bool, *ListNode) {
    if length == 0 {
        return true, node
    } else if length == 1 {
        return true, node.Next
    }
    result, next := isPalindromeHelper(node.Next, length - 2)
    if result == false || next.Val != node.Val {
        return false, nil
    }
    return true, next.Next
}

func getLength(head *ListNode) (length int) {
    for curr := head; curr != nil; curr = curr.Next {length++}
    return
}
