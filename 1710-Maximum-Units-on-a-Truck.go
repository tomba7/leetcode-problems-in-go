/*
 truckSize = 8
 boxTypes = [[3, 10], [6, 5], [4, 7], [2, 9]]
 
 (3, 10), (2, 9), (4, 7), (6, 5)
 truckSize = 8
 */
func maximumUnits(boxTypes [][]int, truckSize int) int {
    sort.Slice(boxTypes, func(i, j int) bool {
        return boxTypes[i][1] > boxTypes[j][1]
    })
    var res int
    for _, boxType := range boxTypes {
        numBoxes := min(truckSize, boxType[0])
        res += numBoxes * boxType[1]
        truckSize -= numBoxes
        if truckSize == 0 {
            break
        }
    }
    return res
}

func min(x, y int) int { if x < y { return x }; return y }

---

// Non-optimal
func maximumUnits(boxTypes [][]int, truckSize int) int {
    sort.Slice(boxTypes, func(i, j int) bool {
        return boxTypes[i][1] > boxTypes[j][1]
    })
    var res int
    var totalBoxes int
    for i := 0; i < len(boxTypes); i++ {
        numBoxes, numUnits := boxTypes[i][0], boxTypes[i][1]
        if totalBoxes + numBoxes < truckSize {
            totalBoxes += numBoxes
            res += numBoxes * numUnits
            continue
        }
        for j := 0; j < numBoxes; j++ {
            totalBoxes++
            res += numUnits
            if totalBoxes == truckSize {
                break
            }
        }
        if totalBoxes == truckSize { break }
    }
    return res
}
