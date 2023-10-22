func construct(grid [][]int) *Node {
    n := len(grid)
    return constructHelper(grid, 0, 0, n-1, n-1)
}

func constructHelper(grid [][]int, top, left, bottom, right int) *Node {
    if top == bottom {
        var val bool
        if grid[top][left] == 1 {
            val = true
        }
        return &Node{IsLeaf: true, Val: val}
    }
    n := (bottom - top + 1)/2
    topLeft     := constructHelper(grid, top, left, top+n-1, left+n-1)
    topRight    := constructHelper(grid, top, left+n, top+n-1, right)
    bottomLeft  := constructHelper(grid, top+n, left, bottom, left+n-1)
    bottomRight := constructHelper(grid, top+n, left+n, bottom, right)
    if topLeft.Val == topRight.Val && topLeft.Val == bottomLeft.Val && topLeft.Val == bottomRight.Val &&
        topLeft.IsLeaf && topRight.IsLeaf && bottomLeft.IsLeaf && bottomRight.IsLeaf {
        return &Node{IsLeaf: true, Val: topLeft.Val}
    }
    fmt.Println("Not Equal")
    return &Node{
        IsLeaf: false, TopLeft: topLeft, TopRight: topRight, BottomLeft: bottomLeft, BottomRight: bottomRight,
    }
}
