/*
 lo = 0, hi = 99
-1[0, 1, 3, 50, 75] 100
                    ^
 
 ["2", "4->49", "51->74", 76-99]
 
 - Walk thru nums
 - If the prev elem is the same, skip
 - If the prev elem is not the same, handle the rance [prev+1, curr-1]
 - Handle edge case when i == n.
 */
func findMissingRanges(nums []int, lower int, upper int) []string {
    var res []string
    prev := lower-1
    for i := 0; i <= len(nums); i++ {
        var curr int
        if i == len(nums) {
            curr = upper + 1
        } else {
            curr = nums[i]
        }
        if curr != prev + 1 {
            start, end := prev + 1, curr - 1
            if start == end {
                res = append(res, fmt.Sprintf("%d", start))
            } else {
                res = append(res, fmt.Sprintf("%d->%d", start, end))
            }
        }
        prev = curr
    }
    return res
}

---

func findMissingRanges(nums []int, lower int, upper int) []string {
    result := []string{}
    prev := lower - 1
    curr := 0
    for i := 0; i <= len(nums); i++ {
        if i == len(nums) {
            curr = upper + 1
        } else {
            curr = nums[i]
        }
        if curr - prev > 1 {
            if prev + 1 == curr - 1 {
                result = append(result, strconv.Itoa(prev + 1))
            } else {
                result = append(result, fmt.Sprintf("%s->%s",
                                strconv.Itoa(prev + 1), strconv.Itoa(curr - 1)))
            }
        }
        prev = curr
    }
    return result
}
