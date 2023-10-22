func closestValue(root *TreeNode, target float64) int {
    var res = root.Val
    for curr := root; curr != nil; {
        if diff(curr.Val, target) < diff(res, target) {
            res = curr.Val
        }
        if target < float64(curr.Val) {
            curr = curr.Left
        } else {
            curr = curr.Right
        }
    }
    return res
}

func diff(val int, target float64) float64 { return math.Abs(float64(val)-target) }

---

func closestValue(root *TreeNode, target float64) int {
    var closestSoFar = math.MaxFloat64
    var res int
    for root != nil {
        currVal := float64(root.Val)
        diff := math.Abs(currVal - target)
        if diff < closestSoFar {
            closestSoFar = diff
            res = root.Val
        }
        if target < currVal {
            root = root.Left
        } else {
            root = root.Right
        }
    }
    return res
}

---

func closestValue(root *TreeNode, target float64) int {
    result := root.Val
    for root != nil {
        if math.Abs(target-float64(root.Val)) < math.Abs(target-float64(result)) {
            result = root.Val
        }
        if target < float64(root.Val) {
            root = root.Left
        } else {
            root = root.Right
        }
    }
    return result
}

---

func closestValue(root *TreeNode, target float64) int {
    var val int
    diff := math.MaxFloat64
    for root != nil {
        if currDiff := math.Abs(target - float64(root.Val)); currDiff < diff {
            diff = currDiff
            val = root.Val
        }
        if target < float64(root.Val) {
            root = root.Left
        } else {
            root = root.Right
        }
    }
    return val
}

---

func closestValue(root *TreeNode, target float64) int {
    var res = root.Val
    closestValueHelper(root, target, &res)
    return res
}

func closestValueHelper(root *TreeNode, target float64, res *int) {
    if root == nil { return }
    if math.Abs(float64(root.Val) - target) < math.Abs(float64(*res)-target) {
        *res = root.Val
    }
    if target < float64(root.Val) {
        closestValueHelper(root.Left, target, res)
    } else {
        closestValueHelper(root.Right, target, res)
    }
}
