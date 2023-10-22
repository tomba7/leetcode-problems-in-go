/*
 - In-order traversal
 - Point prev to dummy node. When we get to the left most node,
   prev.next will point to that node. This is so we can return the
   smallest element in the linked list at the end
 - At each node, we set
   prev.Right = curr
   curr.Left = prev
   prev = curr
 */
func treeToDoublyList(root *Node) *Node {
    if root == nil { return nil }
    dummy := Node{}
    prev := &dummy
    inorder(root, &prev)
    dummy.Right.Left = prev
    prev.Right = dummy.Right
    return dummy.Right
}

func inorder(root *Node, prev **Node) {
    if root == nil {
        return
    }
    inorder(root.Left, prev)
    (*prev).Right = root
    root.Left = *prev
    *prev = root
    inorder(root.Right, prev)
}
