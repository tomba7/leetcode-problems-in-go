type BSTIterator struct {
    stack *list.List
}

func Constructor(root *TreeNode) BSTIterator {
    s := list.New()
    for curr := root; curr != nil; curr = curr.Left {
        s.PushBack(curr)
    }
    return BSTIterator{stack: s}
}

func (it *BSTIterator) Next() int {
    s := it.stack
    curr := s.Remove(s.Back()).(*TreeNode)
    val := curr.Val
    for curr = curr.Right; curr != nil; curr = curr.Left {
        s.PushBack(curr)
    }
    return val
}

func (it *BSTIterator) HasNext() bool {
    return it.stack.Len() != 0
}

---

type BSTIterator struct {
    s []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
    it := BSTIterator{}
    for ; root != nil; root = root.Left {
        it.s = append(it.s, root)
    }
    return it
}

/** @return the next smallest number */
func (it *BSTIterator) Next() int {
    root := it.s[len(it.s)-1]
    it.s = it.s[:len(it.s)-1]
    next := root.Val
    root = root.Right
    for ; root != nil; root = root.Left {
        it.s = append(it.s, root)
    }
    return next
}

/** @return whether we have a next smallest number */
func (it *BSTIterator) HasNext() bool {
    return len(it.s) != 0
}
