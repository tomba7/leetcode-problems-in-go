func minKnightMoves(x int, y int) int {
    if x == 0 && y == 0 { return 0 }
    seen := [605][605]bool{}
    q := list.New()
    q.PushBack([]int{0, 0, 0})
    seen[302][302] = true
    for q.Len() != 0 {
        cell := q.Remove(q.Front()).([]int)
        for _, offset := range getOffsets() {
            nextX, nextY := cell[0] + offset[0], cell[1] + offset[1]
            if seen[nextX+302][nextY+302] { continue }
            seen[nextX+302][nextY+302] = true 
            if nextX == x && nextY == y {
                return cell[2] + 1
            }
            q.PushBack([]int{nextX, nextY, cell[2]+1})
        }
    }
    return -1
}

func getOffsets() [8][2]int {
    offsets := [8][2]int{
        {-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1},
    }
    return offsets
}
