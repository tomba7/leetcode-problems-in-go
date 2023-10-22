/*
  - Set maxHeightSoFar = 0
  - Traverse from right to left
  - If the current bulding's height is taller than the maxHeightSoFar
    add it to the result array and update the maxHeightSoFar to current
  - When done, reverse the results array
*/
func findBuildings(heights []int) []int {
    var maxHeight int
    var res []int
    for i := len(heights)-1; i >= 0; i-- {
        if heights[i] > maxHeight {
            res = append(res, i)
            maxHeight = heights[i]
        }
    }
    for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
        res[i], res[j] = res[j], res[i]
    }
    return res
}

---

/*
  - Traverse from right to left
  - For each element, pop the stack as long as the top
    is lesser than the current element.
  - If the stack is empty (meaning no building is blocking the view)
    add the element to the result
  - Then push the current elemt onto the stack
  - Reverse the res array and return
 */
func findBuildings(heights []int) []int {
    s := list.New()
    var res []int
    for i := len(heights)-1; i >= 0; i-- {
        for s.Len() != 0 && heights[s.Back().Value.(int)] < heights[i] {
            s.Remove(s.Back())
        }
        if s.Len() == 0 {
            res = append(res, i)
        }
        s.PushBack(i)
    }
    reverse(res)
    return res
}

func reverse(nums []int) {
    for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
        nums[i], nums[j] = nums[j], nums[i]
    }
}
