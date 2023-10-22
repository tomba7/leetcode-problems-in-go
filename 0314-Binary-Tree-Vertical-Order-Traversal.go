// BFS
func verticalOrder(root *TreeNode) [][]int {
    if root == nil { return nil }
    table := make(map[int][]int)
    var minCol, maxCol int
    q := list.New()
    q.PushBack(elem{node: root, col: 0})
    for q.Len() != 0 {
        curr := q.Remove(q.Front()).(elem)
        col, node := curr.col, curr.node
        table[col] = append(table[col], node.Val)
        minCol = min(minCol, col)
        maxCol = max(maxCol, col)
        if node.Left != nil {
            q.PushBack(elem{node: node.Left, col: col-1})
        }
        if node.Right != nil {
            q.PushBack(elem{node: node.Right, col: col+1})
        }
    }
    var res [][]int
    for i := minCol; i <= maxCol; i++ {
        res = append(res, table[i])
    }
    return res
}

type elem struct {
    node *TreeNode
    col int
}

func min(x, y int) int { if x < y { return x }; return y }
func max(x, y int) int { if x > y { return x }; return y }

---

// DFS
func verticalOrder(root *TreeNode) [][]int {
    var minCol, maxCol int
    table := make(map[int][][]int)
    dfs(root, 0, 0, &minCol, &maxCol, table)
    var res [][]int
    for i := minCol; i <= maxCol; i++ {
        if len(table[i]) == 0 { continue }
        sort.SliceStable(table[i], func(j, k int) bool {
            return table[i][j][0] < table[i][k][0]
        })
        var sortedColumn []int
        for j := 0; j < len(table[i]); j++ {
            sortedColumn = append(sortedColumn, table[i][j][1])
        }
        res = append(res, sortedColumn)
    }
    return res
}

func dfs(root *TreeNode, row, col int, minCol, maxCol *int, table map[int][][]int) {
    if root == nil { return }
    table[col] = append(table[col], []int{row, root.Val})
    *minCol = min(*minCol, col)
    *maxCol = max(*maxCol, col)
    dfs(root.Left, row+1, col-1, minCol, maxCol, table)
    dfs(root.Right, row+1, col+1, minCol, maxCol, table)
}

func min(x, y int) int { if x < y { return x }; return y }
func max(x, y int) int { if x > y { return x }; return y }

---

type Pair struct {
    col int
    node *TreeNode
}

func verticalOrder(root *TreeNode) [][]int {
    var result [][]int
    var q []Pair
    table := map[int][]int{}
    var minCol, maxCol int = 1, -1
    if root != nil { q = append(q, Pair{0, root}) }
    for len(q) != 0 {
        pair := q[0]
        q = q[1:]
        node, col := pair.node, pair.col
        minCol, maxCol = min(minCol, col), max(maxCol, col)
        table[col] = append(table[col], node.Val)
        if node.Left != nil { q = append(q, Pair{col - 1, node.Left}) }
        if node.Right != nil { q = append(q, Pair{col + 1, node.Right}) }
    }
    for i := minCol; i <= maxCol; i++ { result = append(result, table[i]) }
    return result;
}

func min(x, y int) int { if x < y { return x }; return y }
func max(x, y int) int { if x > y { return x }; return y }
