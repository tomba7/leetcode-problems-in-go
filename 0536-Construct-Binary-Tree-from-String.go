func str2tree(s string) *TreeNode {
    stack := list.New()
    var j int
    for i := 0; i < len(s); i, j = i+1, j+1 {
        ch := s[i]
        if ch == ')' {
            stack.Remove(stack.Back())
        } else {
            for ; i < len(s)-1 && digit(s[i+1]); i++ {}
            val, _ := strconv.Atoi(s[j:i+1])
            newNode := &TreeNode{Val: val}
            if stack.Len() != 0 {
                parent := stack.Back().Value.(*TreeNode)
                if parent.Left == nil {
                    parent.Left = newNode
                } else {
                    parent.Right = newNode;
                }
            }
            stack.PushBack(newNode)
            j = i
        }
        // Skip & ignore '('s
    }
    if stack.Len() == 0 { return nil }
    return stack.Back().Value.(*TreeNode)
}

func digit(ch byte) bool { return '0' <= ch && ch <= '9' }

---

func str2tree(s string) *TreeNode {
    stack := list.New()
    var j int
    for i := 0; i < len(s); i++ {
        ch := s[i]
        if ch == '(' {
            j++
        } else if ch == ')' {
            stack.Remove(stack.Back())
            j++
        } else {
            for ; i < len(s)-1 && digit(s[i+1]); i++ {}
            val, _ := strconv.Atoi(s[j:i+1])
            newNode := &TreeNode{Val: val}
            if stack.Len() != 0 {
                parent := stack.Back().Value.(*TreeNode)
                if parent.Left != nil {
                    parent.Right = newNode;
                } else {
                    parent.Left = newNode
                }
            }
            stack.PushBack(newNode)
            j = i+1
        }
    }
    if stack.Len() == 0 { return nil }
    return stack.Back().Value.(*TreeNode)
}

func digit(ch byte) bool { return '0' <= ch && ch <= '9' }
