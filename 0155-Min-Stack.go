type MinStack struct {
    stack []*Data
}

type Data struct {x, min int}

/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}
}

func min(x, y int) int {if x < y {return x}; return y}

func (s *MinStack) Push(x int)  {
    data := &Data{x: x}
    if len(s.stack) == 0 {
        data.min = x
    } else {
        data.min = min(x, s.GetMin())
    }
    s.stack = append(s.stack, data)
}

func (s *MinStack) Pop()  {
    s.stack = s.stack[:len(s.stack)-1]
}

func (s *MinStack) Top() int {
    return s.stack[len(s.stack)-1].x
}

func (s *MinStack) GetMin() int {
    return s.stack[len(s.stack)-1].min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
