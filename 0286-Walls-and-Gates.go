/*
    _ W G _     3 W G 1
    _ _ _ W     2 2 1 W
    _ W _ W  =  1 W 2 W
    G W _ _     G W 3 4
*/
const (
    wall = -1
    gate = 0
    inf  = 1<<31 - 1
)

var offsets = [][]int{
    {0, 1}, {-1, 0}, {0, -1}, {1, 0},
}

func updateRooms(rooms [][]int, i, j, dist int) {
    m, n := len(rooms), len(rooms[0])
    if i < 0 || i == m || j < 0 || j == n ||
        rooms[i][j] == wall || rooms[i][j] == gate {
        return
    }
    if rooms[i][j] == inf || dist < rooms[i][j] {
        rooms[i][j] = dist
        for _, offset := range offsets {
            updateRooms(rooms, i + offset[0], j + offset[1], dist+1)
        } 
    } 
}

func wallsAndGates(rooms [][]int)  {
    m, n := len(rooms), len(rooms[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if rooms[i][j] == gate {
                for _, offset := range offsets {
                    updateRooms(rooms, i + offset[0], j + offset[1], 1)
                } 
            }
        }
    }
}

---

const (
    wall = -1
    gate = 0
    inf  = 1<<31 - 1
)

var offsets = [][]int{
    {0, 1}, {-1, 0}, {0, -1}, {1, 0},
}

func wallsAndGates(rooms [][]int)  {
    m, n := len(rooms), len(rooms[0])
    q := list.New()
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if rooms[i][j] == gate {
                q.PushBack([]int{i, j})
            }
        }
    }
    for q.Len() != 0 {
        point := q.Remove(q.Front()).([]int)
        row, col := point[0], point[1]
        for _, offset := range offsets {
            i, j := row + offset[0], col + offset[1]
            if i < 0 || i == m || j < 0 || j == n || rooms[i][j] != inf {
                continue
            }
            rooms[i][j] = rooms[row][col] + 1
            q.PushBack([]int{i, j})
        }
    }
}
