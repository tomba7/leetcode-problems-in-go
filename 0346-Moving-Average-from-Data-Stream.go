/* 
  1, 10,  3,   5
  1. 5.5, 4.6, 6.0

  sum = 14
  
  1, 10, 3
- Use a deq, and append numbers as they come in
- Maintain a global sum. Add the incoming numbers here
- If the size equals 'size', remove the front and subtract it from the sum
  Then push back the new elemenent. Add its value to the global sum.
  Then return the sum/size
 */
type MovingAverage struct {
    sum int
    size int
    deque *list.List
}

func Constructor(size int) MovingAverage {
    return MovingAverage{sum: 0, size: size, deque: list.New()}
}

func (a *MovingAverage) Next(val int) float64 {
    if a.deque.Len() == a.size {
        oldVal := a.deque.Remove(a.deque.Front()).(int)
        a.sum -= oldVal
    }
    a.deque.PushBack(val)
    a.sum += val
    return float64(a.sum)/float64(a.deque.Len())
}

---

/*
    size = 3
    [0, 1, 2]
           ^
 */
type MovingAverage struct {
    window []int
    size, insert, sum int
}

func Constructor(size int) MovingAverage {
    return MovingAverage{window: make([]int, size)}
}

func (a *MovingAverage) Next(val int) float64 {
    if a.size < len(a.window) {
        a.size++
    }
    a.sum -= a.window[a.insert]
    a.sum += val
    a.window[a.insert] = val
    a.insert = (a.insert + 1) % len(a.window)
    return float64(a.sum)/float64(a.size)
}

---

type MovingAverage struct {
    size, sum int
    window *list.List
}

func Constructor(size int) MovingAverage {
    return MovingAverage{
        size: size, window: list.New()
    }
}

func (ma *MovingAverage) Next(val int) float64 {
    if ma.window.Len() == ma.size {
        ma.sum -= ma.window.Front().Value.(int)
        ma.window.Remove(ma.window.Front())
    }
    ma.sum += val
    ma.window.PushBack(val)
    return float64(ma.sum)/float64(ma.window.Len())
}
