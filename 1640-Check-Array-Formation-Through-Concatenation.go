/*
    arr = [91,4,64,78], pieces = [[78],[4,64],[91]]
    
    t[78] = [78]
    t[4]  = [4,64]
    t[91] = [91]

    [91,4,64,78]
             ^
 */
func canFormArray(arr []int, pieces [][]int) bool {
    table := make(map[int][]int)
    for _, piece := range pieces {
        table[piece[0]] = piece
    }
    var i int
    for i < len(arr) {
        piece, exist := table[arr[i]]
        if !exist {
            return false
        }
        for j := 0; j < len(piece); j++ {
            if piece[j] != arr[i] {
                return false
            }
            i++
        }
    }
    return true
}
