/*
  Return:
    - 0 indexed
    - Point with the smallest distance from curr loc with same x or y
    - If multiple, return point with smallest index
    - If none, return -1

    x = 3, y = 4
    
    [1,2],[3,1],[2,4],[2,3],[4,4]
      *     c.    c.    c.    c
                              i
            
      dist =  (x1 - x2)^2 - (y1 - y2)^2
      minDistSoFar = 9
      res = 1
  Algo
  - Maintain a running min
  - For each points
    * ignore it if it does not contain x or y
    * compute the current distance
    * if dist < res: update res = dist
    
 */
func nearestValidPoint(x int, y int, points [][]int) int {
    var res = -1
    var minDistSoFar = math.MaxInt32 
    distance := func(point []int) int {
        d1, d2 := x - point[0], y - point[1]
        return d1*d1 + d2*d2
    }
    for i, point := range points {
        if x == point[0] || y == point[1] {
            currDistance := distance(point)
            if currDistance < minDistSoFar {
                minDistSoFar = currDistance
                res = i
            }
        }
    }
    return res 
}

---

func nearestValidPoint(x int, y int, points [][]int) int {
    var res = -1
    var minDistSoFar = math.MaxInt32 
    for i, point := range points {
        dx, dy := abs(x-point[0]), abs(y-point[1])
        dist := dx*dx + dy*dy
        if dx*dy == 0 && dist < minDistSoFar {
            minDistSoFar = dist
            res = i
        }
    }
    return res 
}
