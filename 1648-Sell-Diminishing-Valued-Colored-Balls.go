func maxProfit(inventory []int, orders int) int {
    n := len(inventory)
    var mod int = 1e9+7
    sort.Slice(inventory, func(i, j int) bool {
        return inventory[i] > inventory[j]
    })
    numColors := 1
    var res int
    for i := 0; i < n; i++ {
        curr := inventory[i]
        var next int
        if i != n - 1 {
            next = inventory[i+1]
        }
        numBalls := min(curr-next, orders/numColors)
        orders -= numBalls * numColors
        res = (res + (curr*(curr+1)-(curr-numBalls)*(curr-numBalls+1))/2*numColors) % mod
        if curr-next > numBalls {
            res = (res + orders * (curr-numBalls)) % mod
            break
        }
        numColors++
    }
    return res
}

func min(x, y int) int { if x < y { return x }; return y }

---

// Brute force (TLE)
func maxProfit(inventory []int, orders int) int {
    h := &Heap{}
    for _, ball := range inventory {
        heap.Push(h, ball)
    }
    var res int
    for h.Len() != 0 && orders > 0 {
        ball := heap.Pop(h).(int)
        res += ball
        heap.Push(h, ball-1)
        orders--
    }
    return res%(1e9+7)
}

type Heap []int
func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] > h[j] }
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *Heap) Pop() interface{} { x := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return x }

---

/*
    [2,8,4,10,6]
    
    [10,8,6,4,2]    orders = 20
            i
    ------------------------------------------------------ 
    i = 3
    curr      = 4
    next      = 2
    numColors = 4
    numBalls  = 2
    orders    = orders - numBalls * numColors
              = 8 - 2 * 4
              = 0
    res       = res + [n1(n1+1)/2 - n2(n2+1)/2]*numColors
              = 82 + [4(5)/2 - 2(3)/2] * 4
              = 82 + 28
              = 110
    ------------------------------------------------------ 
    i = 2
    curr      = 6
    next      = 4
    numColors = 3
    numBalls  = 2
    orders    = orders - numBalls * numColors
              = 14 - 2 * 3
              = 8
    res       = res + [n1(n1+1)/2 - n2(n2+1)/2]*numColors
              = 49 + [6(7)/2 - 4(5)/2] * 3
              = 49 + 33
              = 82
    ------------------------------------------------------
    i = 1
    curr      = 8
    next      = 6
    numColors = 2
    numBalls  = 2
    orders    = orders - numBalls * numColors
              = 18 - 2 * 2
              = 14
    res       = res + [n1(n1+1)/2 - n2(n2+1)/2]*numColors
              = 19 + [8(8+1)/2 - 6(6+1)/2]*2
              = 19 + 30
              = 49
    ------------------------------------------------------
    i = 0
    curr      = 10
    next      = 8
    numColors = 1
    numBalls  = 2
    orders    = orders - numBalls * numColors
              = 20 - 2 x 1
              = 18
    res       = res + [n1(n1+1)/2 - n2(n2+1)/2]*numColors
              = 19
 */
