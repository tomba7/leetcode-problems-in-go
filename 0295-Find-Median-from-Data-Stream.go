type MedianFinder struct {
    lower, upper *Heap
}

func Constructor() MedianFinder {
    return MedianFinder{
        lower: NewHeap(true), upper: NewHeap(false),
    }
}

func (m *MedianFinder) AddNum(num int)  {
    lower, upper := m.lower, m.upper
    heap.Push(lower, num)
    heap.Push(upper, heap.Pop(lower).(int))
    if lower.Len() < upper.Len() {
        heap.Push(lower, heap.Pop(upper).(int))
    }
}

func (m *MedianFinder) FindMedian() float64 {
    var median float64
    size := m.lower.Len() + m.upper.Len()
    if size & 1 == 0 {
        median = float64(m.lower.Top() + m.upper.Top())/2.0
    } else {
        median = float64(m.lower.Top())
    }
    return median
}

type Heap struct {
    min bool
    nums []int
}

func NewHeap(min bool) *Heap {
    return &Heap{min: min}
}

func (h *Heap) Top() int { return h.nums[0] }
func (h *Heap) Len() int { return len(h.nums) }
func (h *Heap) Less(i, j int) bool {
    if h.min {
        return h.nums[i] < h.nums[j]
    }
    return h.nums[i] > h.nums[j]
}
func (h *Heap) Swap(i, j int) {
    h.nums[i], h.nums[j] = h.nums[j], h.nums[i]
}
func (h *Heap) Push(x interface{}) {
    h.nums = append(h.nums, x.(int))
}
func (h *Heap) Pop() interface{} {
    x := h.nums[len(h.nums)-1]
    h.nums = h.nums[:len(h.nums)-1]
    return x
}

---

/*
 - split the data into half at any point in time
 - lower(max heap) + upper(min heap)
 - sizeof(lower) + 1 = sizeof(upper)
 - To get the median
   odd:
    get from lower
   even:
    get from both, add, divide/2
    
 - To add numbers
   * If len(lower) == len(upper)
     - if lower.Len() == 0 || x <= upper.Peek()
       lower.Push(x)
     - else
       y = upper.Pop() and lower.Push(y)
       upper.Push(x)
       
   * Else
     - if x < lower.Peek()
       y = lower.Pop(), upper.Push(y) 
       lower.Push(x)
     - else
       upper.Push(x)
       
   0,1      2,3
   
   
   3,5,8    9,10
 */
type MedianFinder struct {
    lower, upper *Heap
}

func Constructor() MedianFinder {
    return MedianFinder{
        lower: NewHeap(false), upper: NewHeap(true),
    }
}

func (f *MedianFinder) AddNum(num int)  {
    lower, upper := f.lower, f.upper
    if lower.Len() == upper.Len() {
        if lower.Len() == 0 || num <= upper.Peek() {
            heap.Push(lower, num)
        } else {
            top := heap.Pop(upper).(int)
            heap.Push(lower, top)
            heap.Push(upper, num)
        }
    } else {
        if num < lower.Peek() {
            top := heap.Pop(lower).(int)
            heap.Push(upper, top)
            heap.Push(lower, num)
        } else {
            heap.Push(upper, num)
        }
    }
}

func (f *MedianFinder) FindMedian() float64 {
    lower, upper := f.lower, f.upper
    if lower.Len() == upper.Len() {
        x, y := float64(lower.Peek()), float64(upper.Peek())
        return (x + y)/2
    }
    return float64(lower.Peek())
}

type Heap struct {
    min bool
    nums []int
}

func NewHeap(min bool) *Heap {
    return &Heap{min: min}
}

func (h *Heap) Peek() int { return h.nums[0] }
func (h *Heap) Len() int { return len(h.nums) }
func (h *Heap) Less(i, j int) bool {
    if h.min {
        return h.nums[i] < h.nums[j]
    }
    return h.nums[i] > h.nums[j]
}
func (h *Heap) Swap(i, j int) {
    h.nums[i], h.nums[j] = h.nums[j], h.nums[i]
}
func (h *Heap) Push(x interface{}) {
    h.nums = append(h.nums, x.(int))
}
func (h *Heap) Pop() interface{} {
    x := h.nums[len(h.nums)-1]
    h.nums = h.nums[:len(h.nums)-1]
    return x
}
