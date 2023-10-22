/*
 - State (pr, pc, br, bc)
 - Begin from the start cell and enqueue all valid next positions
 - Set the result to MAX_INT
 - Mark a state as seen with a dist of 0
 - If we find the box in the the player's next cell we start moving
   the box's postion in the same direction and update its distance by 1
   * If the new state already exist AND the existing distance is less than
     the curr + 1, then continue. This is because the existing distance is
     already optimal, so we dont want to update it with a less optimal
     distance.
   * Else, update the distance of the new state to curr + 1
   * Then enqueue
 - If we do not find the box in the next cell, we simply check if the 
   cell has been visited or not.
   * If the state exist (has been visited) AND if the existing distance
     is more optimal (lesser) than curr + 1, then ignore.
 - When we dequeue a cell from the queue, if the dist of the cell
   is greater or equal to the current result, ignore it
 - If the curr cell is the target cell, update the min(res, dist[curr][cell])
*/
func minPushBox(grid [][]byte) int {
    q := list.New()
    box, target, keeper := getBoxTargetAndStoreKeeper(grid)
    state := &State{box: box, keeper: keeper}
    dist := map[string]int{key(state): 0}
    q.PushBack(state)
    res := math.MaxInt32
    for q.Len() != 0 {
        curr := q.Remove(q.Front()).(*State)
        if dist[key(curr)] >= res { continue }
        currBox, currKeeper := curr.box, curr.keeper
        if currBox[0] == target[0] && currBox[1] == target[1] {
            res = min(res, dist[key(curr)])
            continue
        }
        for _, offset := range [][]int{{0,-1}, {0,1}, {-1,0}, {1,0}} {
            newKeeper := []int{currKeeper[0] + offset[0], currKeeper[1] + offset[1]}
            if !valid(newKeeper, grid) { continue }
            if newKeeper[0] == currBox[0] && newKeeper[1] == currBox[1] {
                newBox := []int{currBox[0] + offset[0], currBox[1] + offset[1]}
                if !valid(newBox, grid) { continue }
                next := &State{box: newBox, keeper: newKeeper}
                if di, exist := dist[key(next)]; exist && di <= dist[key(curr)] + 1 {
                    continue
                }
                dist[key(next)] = dist[key(curr)] + 1
                q.PushBack(next)
            } else {
                next := &State{box: currBox, keeper: newKeeper}
                if di, exist := dist[key(next)]; exist && di <= dist[key(curr)] {
                    continue
                }
                dist[key(next)] = dist[key(curr)]
                q.PushBack(next)
            }
        }
    }
    if res == math.MaxInt32 { return -1 }
    return res
}

type State struct {
    box, keeper []int
}

func getBoxTargetAndStoreKeeper(grid [][]byte) ([]int, []int, []int) {
    m, n := len(grid), len(grid[0])
    var box, target, keeper []int
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 'B' {
                box = []int{i, j}
            } else if grid[i][j] == 'T' {
                target = []int{i, j}
            } else if grid[i][j] == 'S' {
                keeper = []int{i, j}
            }
        }
    }
    return box, target, keeper
}

func key(state *State) string {
    return fmt.Sprintf("%d,%d,%d,%d",
        state.box[0], state.box[1],
        state.keeper[0], state.keeper[1])
}

func valid(cell []int, grid [][]byte) bool {
    m, n, r, c := len(grid), len(grid[0]), cell[0], cell[1]
    return 0 <= r && r < m && 0 <= c && c < n && grid[r][c] != '#'
}

func min(x, y int) int { if x < y { return x }; return y }

---

// This does not work yet. Directly translated from the following link
// https://leetcode.com/problems/minimum-moves-to-move-a-box-to-their-target-location/discuss/693918/Python-BFS-*-BFS-or-130ms-or-beats-95-or-Explained-and-Commented
func minPushBox(grid [][]byte) int {
    box, target, keeper := getBoxTargetAndStoreKeeper(grid)
    q := list.New()
    q.PushBack(&elem{dist: 0, box: box, keeper: keeper})
    seen := map[string]bool{key2(box, keeper): true}
    for q.Len() != 0 {
        e := q.Remove(q.Front()).(*elem)
        dist, box, keeper := e.dist, e.box, e.keeper
        if box[0] == target[0] && box[1] == target[1] {
            return dist
        }
        boxOffset    := [][]int{{0,-1}, {0,1}, {-1,0}, {1,0}}
        keeperOffset := [][]int{{0,1}, {0,-1}, {1,0}, {-1,0}}
        for i := 0; i < len(boxOffset); i++ {
            newBox := []int{box[0] + boxOffset[i][0], box[1] + boxOffset[i][1]}
            newKeeper := []int{box[0] + keeperOffset[i][0], box[1] + keeperOffset[i][1]}
            if valid(newBox[0], newBox[1], grid) && !seen[key2(newBox, box)] {
                if valid(newKeeper[0], newKeeper[1], grid) && check(keeper, newKeeper, box, grid) {
                    seen[key2(newBox, box)] = true
                    q.PushBack(&elem{dist: dist+1, box:newBox, keeper: box})
                }
            }
        }
    }
    return -1
}

type elem struct {
    dist int
    box, keeper []int
}

func check(curr, dest, box []int, grid [][]byte) bool {
    q := list.New()
    q.PushBack(curr)
    seen := make(map[string]bool)
    for q.Len() != 0 {
        pos := q.Remove(q.Front()).([]int)
        if pos[0] == dest[0] && pos[1] == dest[1] {
            return true
        }
        for _, offset := range [][]int{{0,-1}, {0,1}, {-1,0}, {1,0}} {
            x, y := pos[0] + offset[0], pos[1] + offset[1]
            if valid(x, y, grid) && !seen[key1(x, y)] && x != box[0] && y != box[1] {
                seen[key1(x, y)] = true
                q.PushBack([]int{x, y})
            }
        }
    }
    return false
}

func key1(x, y int) string { return fmt.Sprintf("%d,%d", x, y) }
func key2(x, y []int) string { return fmt.Sprintf("%d,%d,%d,%d", x[0], x[1], y[0], y[1]) }

func getBoxTargetAndStoreKeeper(grid [][]byte) ([]int, []int, []int) {
    m, n := len(grid), len(grid[0])
    var box, target, keeper []int
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 'B' {
                box = []int{i, j}
            } else if grid[i][j] == 'T' {
                target = []int{i, j}
            } else if grid[i][j] == 'S' {
                keeper = []int{i, j}
            }
        }
    }
    return box, target, keeper
}

func valid(r, c int, grid [][]byte) bool {
    m, n := len(grid), len(grid[0])
    return 0 <= r && r < m && 0 <= c && c < n && grid[r][c] != '#'
}

func min(x, y int) int { if x < y { return x }; return y }
