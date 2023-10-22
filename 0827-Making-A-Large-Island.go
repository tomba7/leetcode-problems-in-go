/*
    1,0     1,1     1,1
    0,1     1,0     1,1
    
    1,0,1,1    2,0,3,3
    0,1,0,1    0,4,0,3
    1,1,1,0    4,4,4,0
    0,1,1,1    0,4,4,4
    
    2: 1
    3: 3
    4: 7
    
Brute Force
- From every 0, start a DFS: O(n^2)
- After every DFS keep track of the max island size
- Return the max island size

Optimal
- Find out the size of each island, assign them an id, then store the id => area mappings
- Go thru all the 0's, get the neigbors, then compute the area. Track the max and return
 */
func largestIsland(grid [][]int) int {
    n := len(grid)
    area := make(map[int]int)
    var res int
    id := 2
    for r := 0; r < n; r++ {
        for c := 0; c < n; c++ {
            if grid[r][c] == 1 {
                area[id] = dfs(grid, r, c, id)
                res = max(res, area[id])
                id++
            }
        }
    }
    for r := 0; r < n; r++ {
        for c := 0; c < n; c++ {
            if grid[r][c] != 0 { continue }
            total := 1
            seen := map[int]bool{}
            for _, neighbor := range neighbors(r, c, len(grid)) {
                i, j := neighbor[0], neighbor[1]
                id := grid[i][j]
                if seen[id] { continue }
                seen[id] = true
                total += area[id]
            }
            res = max(res, total)
        }
    }
    return res
}

func dfs(grid [][]int, r, c, id int) int {
    grid[r][c] = id
    var total int
    for _, neighbor := range neighbors(r, c, len(grid)) {
        i, j := neighbor[0], neighbor[1]
        if grid[i][j] == 1 {
            total += dfs(grid, i, j, id)
        }
    }
    return total + 1
}

func neighbors(r, c, n int) [][]int {
    var res [][]int
    if valid(r, c+1, n) { res = append(res, []int{r, c+1}) }
    if valid(r+1, c, n) { res = append(res, []int{r+1, c}) }
    if valid(r, c-1, n) { res = append(res, []int{r, c-1}) }
    if valid(r-1, c, n) { res = append(res, []int{r-1, c}) }
    return res
}

func valid(r, c, n int) bool { return 0 <= r && r < n && 0 <= c && c < n }

func max(x, y int) int { if x > y { return x }; return y }

---

func largestIsland(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    seen := make([][]int, m)
    for i := range seen {
        seen[i] = make([]int, n)
    }
    area := map[int]int{}
    id := 1
    noZero := true
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 0 || seen[i][j] != 0 {
                if grid[i][j] == 0 { noZero = false }
                continue
            }
            currArea := 0
            dfs(i, j, id, grid, seen, area, &currArea)
            area[id] = currArea
            id++
        }
    }
    if noZero { return m * n }
    var result int
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] != 0 {
                continue
            }
            neighbors := map[int]struct{}{}
            currArea := 1
            for _, offset := range offsets {
                x, y := i + offset[0], j + offset[1]
                if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 { continue }
                id = seen[x][y]
                neighbors[id] = struct{}{}
            }
            for id := range neighbors {
                currArea += area[id]
            }
            result = max(result, currArea)
        }
    }
    return result
}

func dfs(i, j, id int, grid, seen [][]int, area map[int]int, currArea *int) {
    m, n := len(grid), len(grid[0])
    if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 || seen[i][j] != 0 {
        return
    }
    seen[i][j] = id
    *currArea++
    for _, offset := range offsets {
        x, y := i + offset[0], j + offset[1]
        dfs(x, y, id, grid, seen, area, currArea)
    }
}

var offsets = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
func max(x, y int) int { if x > y { return x }; return y }

---

func largestIsland(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    seen := make([][]int, m)
    for i := range seen {
        seen[i] = make([]int, n)
    }
    area := map[int]int{}
    id := 1
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 0 || seen[i][j] != 0 {
                //if grid[i][j] == 0 { noZero = false }
                continue
            }
            currArea := 0
            dfs(i, j, id, grid, seen, area, &currArea)
            area[id] = currArea
            id++
        }
    }
    var result int
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] != 0 {
                continue
            }
            neighbors := map[int]struct{}{}
            currArea := 1
            for _, offset := range offsets {
                x, y := i + offset[0], j + offset[1]
                if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 { continue }
                id = seen[x][y]
                neighbors[id] = struct{}{}
            }
            for id := range neighbors {
                currArea += area[id]
            }
            result = max(result, currArea)
        }
    }
    if result == 0 { return m * n }
    return result
}

func dfs(i, j, id int, grid, seen [][]int, area map[int]int, currArea *int) {
    m, n := len(grid), len(grid[0])
    if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 || seen[i][j] != 0 {
        return
    }
    seen[i][j] = id
    *currArea++
    for _, offset := range offsets {
        x, y := i + offset[0], j + offset[1]
        dfs(x, y, id, grid, seen, area, currArea)
    }
}

var offsets = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
func max(x, y int) int { if x > y { return x }; return y }

---

// Brute force (TLE
func largestIsland(grid [][]int) int {
    n := len(grid)
    var res int
    noZeroes := true
    for r := 0; r < n; r++ {
        for c := 0; c < n; c++ {
            if grid[r][c] != 0 { continue }
            noZeroes = false
            grid[r][c] = 1
            seen := newSeen(n)
            res = max(res, dfs(grid, seen, r, c))
            grid[r][c] = 0
        }
    }
    if noZeroes { return n * n }
    return res
}

func dfs(grid [][]int, seen [][]bool, r int, c int) int {
    n := len(grid)
    if r < 0 || r == n || c < 0 || c == n || grid[r][c] == 0 || seen[r][c] { return 0 }
    seen[r][c] = true
    return 1 + dfs(grid, seen, r, c+1) + dfs(grid, seen, r+1, c) +
               dfs(grid, seen, r, c-1) + dfs(grid, seen, r-1, c)
}

func newSeen(n int) [][]bool {
    seen := make([][]bool, n)
    for i := range seen { seen[i] = make([]bool, n) }
    return seen
}

func max(x, y int) int { if x > y { return x }; return y }

