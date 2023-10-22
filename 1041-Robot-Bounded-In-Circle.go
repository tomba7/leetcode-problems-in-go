func isRobotBounded(instructions string) bool {
    directions := [][]int{
        {-1,  0},   // 0: North
        { 0,  1},   // 1: East
        { 1,  0},   // 2: South
        { 0, -1},   // 3: West
    }
    var r, c, i int
    for _, inst := range instructions {
        if inst == 'R' {
            i = (i + 1) % 4
        } else if inst == 'L' {
            i = (i + 3) % 4
        } else {
            r, c = r + directions[i][0], c + directions[i][1]
        }
    }
    return (r == 0 && c == 0) || i > 0
}
