func kClosest(points [][]int, k int) [][]int {
    h := &Heap{}
    for _, point := range points {
        if h.Len() < k {
            heap.Push(h, point)
        } else if dist(point) < dist(h.Peek()) {
            heap.Pop(h)
            heap.Push(h, point)
        }
    }
    var res [][]int
    for h.Len() != 0 {
        res = append(res, heap.Pop(h).([]int))
    }
    return res
}

func dist(point []int) int { return point[0]*point[0] + point[1]*point[1] }

type Heap [][]int

func (h Heap) Peek() []int { return h[0] }
func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool { return dist(h[i]) > dist(h[j]) }
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *Heap) Pop() interface{} { n := len(*h); x := (*h)[n-1]; *h = (*h)[:n-1]; return x }

---

func kClosest(points [][]int, k int) [][]int {
    n := len(points)
    rand.Seed(time.Now().UnixNano())
    lo, hi := 0, n-1
    k--
    for lo <= hi {
        mid := partition(points, lo, hi)
        if mid < k {
            lo = mid + 1
        } else if mid > k {
            hi = mid - 1
        } else {
            return points[:k+1]
        }
    }
    return nil
}

func partition(points [][]int, lo, hi int) int {
    i, j, k := lo, lo, hi
    pivotIndex := lo + rand.Intn(hi-lo+1)
    pivot := points[pivotIndex]
    for j <= k {
        if dist(points[j]) == dist(pivot) {
            j++
        } else if dist(points[j]) < dist(pivot) {
            points[i], points[j] = points[j], points[i]
            i++
            j++
        } else {
            points[j], points[k] = points[k], points[j]
            k--
        }
    }
    return i
}

func dist(point []int) int {
    return point[0]*point[0] + point[1]*point[1]
}

---

/*
    sqrt(x^2 + y^2) = x^2 + y^2

    - Quick select to partition
    - k--
    - In an outer loop [lo=0, hi=len(points)-1], partition(lo, hi) the points array
    - if the pivotIndex equals k, return points[:pivotIndex+1]
    - else if k < pivotIndex, hi=pivotIndex-1
    - else if k > pivotIndex, lo=pivotIndex+1 and k = k - pivotIndex
    k = 6
    
    [3, 12, 2, 8, 15, 27, 17, 19]
                  *
                  pi
 */
func kClosest(points [][]int, k int) [][]int {
    k--
    for lo, hi := 0, len(points)-1; lo <= hi; {
        indx := partition(points, lo, hi)
        if k == indx {
            return points[:indx+1]
        } else if k < indx {
            hi = indx - 1
        } else {
            lo = indx + 1
        }
    }
    return nil
}

func partition(points [][]int, lo, hi int) int {
    indx := lo + rand.Intn(hi-lo+1)
    pivot := points[indx]
    swap(points, indx, hi)
    i := lo
    for j := lo; j < hi; j++ {
        if dist(points[j]) < dist(pivot) {
            swap(points, i, j)
            i++
        }
    }
    swap(points, i, hi)
    return i
}

func dist(point []int) int { return point[0]*point[0] + point[1]*point[1] }
func swap(points [][]int, i, j int) { points[i], points[j] = points[j], points[i] }

---

func kClosest(points [][]int, k int) [][]int {
    list := make([]*Point, len(points))
    for i := 0; i < len(list); i++ {
        x, y := points[i][0], points[i][1]
        list[i] = &Point{
            dist: math.Sqrt(float64((x*x)+(y*y))),
            index: i,
        }
    }
    quickSelect(list, k)
    res := make([][]int, k)
    for i := 0; i < int(k); i++ {
        res[i] = points[list[i].index]
    }
    return res
}

type Point struct {
    dist float64
    index int
}

func quickSelect(list []*Point, k int) {
    lo, hi := 0, len(list)-1
    for lo <= hi {
        mid := partition(list, lo, hi)
        if mid == k {
            return
        } else if mid < k {
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }
}

func partition(list []*Point, lo, hi int) int {
    i, j, k := lo, lo, hi
    pivot := list[lo].dist
    for j <= k {
        if list[j].dist < pivot {
            list[i], list[j] = list[j], list[i]
            i, j = i+1, j+1
        } else if list[j].dist > pivot {
            list[j], list[k] = list[k], list[j]
            k--
        } else {
            j++
        }
    }
    return i
}

---

func kClosest(points [][]int, k int) [][]int {
    h := &maxHeap{}
    for i := 0; i < len(points); i++ {
        x, y := points[i][0], points[i][1]
        heap.Push(h, Point{
            index: i,
            dist: math.Sqrt(float64(x*x + y*y)),
        })
        if h.Len() > k { heap.Pop(h) }
    }
    var res [][]int
    for _, p := range *h {
        res = append(res, points[p.index])
    }
    return res
}

type Point struct {
    index int
    dist float64
}

type maxHeap []Point

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i].dist >= h[j].dist }
func (h maxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(val interface{}) {
    *h = append(*h, val.(Point))
}
func (h *maxHeap) Pop() interface{} {
    val := (*h)[h.Len()-1]
    *h = (*h)[:h.Len()-1]
    return val
}
