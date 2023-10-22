/*
 Overlapping
            1          5
    first   ------------
                 3           7
    second       -------------
                 |     |
                 start end
    
 Non-overlapping
            1          5
    first   ------------ 
                             7          11
    second                   -------------
                       |     | 
                       end   start
 */
func intervalIntersection(first, second [][]int) [][]int {
    var res [][]int
    for i, j := 0, 0; i < len(first) && j < len(second); {
        start := max(first[i][0], second[j][0])
        end := min(first[i][1], second[j][1])
        if start <= end {
            res = append(res, []int{start, end})
        }
        if first[i][1] < second[j][1] {
            i++
        } else {
            j++
        }
    }
    return res
}

func min(x, y int) int { if x < y { return x }; return y }
func max(x, y int) int { if x > y { return x }; return y }

---

func intervalIntersection(first, second [][]int) [][]int {
    var res [][]int
    for i, j := 0, 0; i < len(first) && j < len(second); {
        if first[i][0] < second[j][0] {
            if overlaps(first[i], second[j]) {
                res = append(res, []int{
                    second[j][0], min(first[i][1], second[j][1]),
                })
            }
        } else {
            if overlaps(second[j], first[i]) {
                res = append(res, []int{
                    first[i][0], min(first[i][1], second[j][1]),
                })
            }
        }
        if first[i][1] < second[j][1] {
            i++
        } else {
            j++
        }
    }
    return res
}

func overlaps(first, second []int) bool { return second[0] <= first[1] }
func min(x, y int) int { if x < y { return x }; return y }
