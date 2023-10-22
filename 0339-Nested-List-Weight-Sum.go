func depthSum(nestedList []*NestedInteger) int {
    return depthSumHelper(nestedList, 1)
}

func depthSumHelper(nestedList []*NestedInteger, level int) int {
    var sum int
    for _, elem := range nestedList {
        if elem.IsInteger() {
            sum += elem.GetInteger() * level
        } else {
            sum += depthSumHelper(elem.GetList(), level+1)
        }
    }
    return sum
}

---

func depthSum(nestedList []*NestedInteger) int {
    q := list.New()
    for _, elem := range nestedList {
        q.PushBack(elem)
    }
    var sum int
    level := 1
    for q.Len() != 0 {
        n := q.Len()
        for i := 0; i < n; i++ {
            curr := q.Remove(q.Front()).(*NestedInteger)
            if curr.IsInteger() {
                sum += curr.GetInteger() * level
            } else {
                for _, elem := range curr.GetList() {
                    q.PushBack(elem)
                }
            }
        }
        level++
    }
    return sum
}
