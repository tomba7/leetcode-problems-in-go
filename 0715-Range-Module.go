type RangeModule struct {
    ranges []int
}

func Constructor() RangeModule {
    return RangeModule{}
}

func (m *RangeModule) AddRange(left int, right int)  {
    i, j := findLeft(m.ranges, left), findRight(m.ranges, right)
    var subRange []int
    if i % 2 == 0 { subRange = append(subRange, left) }
    if j % 2 == 0 { subRange = append(subRange, right) }
    m.ranges = insert(m.ranges, i, j, subRange...)
}

func (m *RangeModule) QueryRange(left int, right int) bool {
    i, j := findRight(m.ranges, left), findLeft(m.ranges, right)
    return i == j && i % 2 == 1
}

func (m *RangeModule) RemoveRange(left int, right int)  {
    i, j := findLeft(m.ranges, left), findRight(m.ranges, right)
    var subRange []int
    if i % 2 == 1 { subRange = append(subRange, left) }
    if j % 2 == 1 { subRange = append(subRange, right) }
    m.ranges = insert(m.ranges, i, j, subRange...)
}

func insert(nums []int, i, j int, slice... int) []int {
    slice = append(slice, nums[j:]...)
    return append(nums[:i], slice...)
}

func findLeft(nums []int, left int) int {
    return sort.Search(len(nums), func (i int) bool {
        return nums[i] >= left
    })
}

func findRight(nums []int, right int) int {
    return sort.Search(len(nums), func (i int) bool {
        return nums[i] > right
    })
}

---

type RangeModule []int

func Constructor() RangeModule { return RangeModule{} }

func (m *RangeModule) AddRange(left int, right int)  {
    i, j := m.findLeft(left), m.findRight(right)
    var subRange []int
    if i % 2 == 0 { subRange = append(subRange, left) }
    if j % 2 == 0 { subRange = append(subRange, right) }
    m.insert(i, j, subRange...)
}

func (m *RangeModule) QueryRange(left int, right int) bool {
    i, j := m.findRight(left), m.findLeft(right)
    return i == j && i % 2 == 1
}

func (m *RangeModule) RemoveRange(left int, right int)  {
    i, j := m.findLeft(left), m.findRight(right)
    var subRange []int
    if i % 2 == 1 { subRange = append(subRange, left) }
    if j % 2 == 1 { subRange = append(subRange, right) }
    m.insert(i, j, subRange...)
}

func (m *RangeModule) insert(i, j int, slice... int) {
    slice = append(slice, (*m)[j:]...)
    *m = append((*m)[:i], slice...)
}

func (m *RangeModule) findLeft(left int) int {
    return sort.Search(len(*m), func (i int) bool {
        return (*m)[i] >= left
    })
}

func (m *RangeModule) findRight(right int) int {
    return sort.Search(len(*m), func (i int) bool {
        return (*m)[i] > right
    })
}
