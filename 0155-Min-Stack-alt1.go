type MinStack struct {
    stack []int
    helper []*Helper
}

type Helper struct {
    val, count int
}

/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}
}

func (s *MinStack) Push(x int)  {
    if len(s.stack) == 0 || x < s.helper[len(s.helper)-1].val{
        s.helper = append(s.helper, &Helper{val: x})
    } else if x == s.helper[len(s.helper)-1].val {
        s.helper[len(s.helper)-1].count++
    }
    s.stack = append(s.stack, x)
}

func (s *MinStack) Pop()  {
    if s.helper[len(s.helper)-1].val != s.stack[len(s.stack)-1] {
        s.stack = s.stack[:len(s.stack)-1]
        return
    }
    if s.helper[len(s.helper)-1].count != 0 {
        s.helper[len(s.helper)-1].count--
    } else {
        s.helper = s.helper[:len(s.helper)-1]
    }
    s.stack = s.stack[:len(s.stack)-1]
}

func (s *MinStack) Top() int {
    return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
    return s.helper[len(s.helper)-1].val
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
