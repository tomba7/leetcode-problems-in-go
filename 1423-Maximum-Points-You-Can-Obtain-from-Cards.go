func maxScore(cardPoints []int, k int) int {
    n := len(cardPoints)
    left, right := make([]int, k+1), make([]int, k+1)
    for i := 0; i < k; i++ {
        left[i+1] = left[i] + cardPoints[i]
        right[i+1] = right[i] + cardPoints[n-1-i]
    }
    var res int
    for i := 0; i <= k; i++ {
        res = max(res, left[i] + right[k-i])
    }
    return res
}

func max(x, y int) int { if x > y { return x }; return y }

---

func maxScore(cardPoints []int, k int) int {
    n := len(cardPoints)
    var left, right, res int
    for i := 0; i < k; i++ {
        left += cardPoints[i]
    }
    res = left
    for i := 0; i <= k; i++ {
        res = max(res, left + right)
        if i < k {
            left -= cardPoints[k-1-i]
            right += cardPoints[n-1-i]
        }
    }
    return res
}

func max(x, y int) int { if x > y { return x }; return y }

---

func maxScore(cardPoints []int, k int) int {
    n := len(cardPoints)
    subArrayLength := n - k
    var startIndex, totalScore int
    for _, point := range cardPoints {
        totalScore += point
    }
    minSubArrayScore := totalScore
    if k == n { return totalScore }
    var currScore int
    for i := 0; i < n; i++ {
        currScore += cardPoints[i]
        currSubArrayLength := i - startIndex + 1
        if currSubArrayLength == subArrayLength {
            minSubArrayScore = min(minSubArrayScore, currScore)
            currScore -= cardPoints[startIndex]
            startIndex++
        }
    }
    return totalScore - minSubArrayScore
}

func min(x, y int) int { if x < y { return x }; return y }

---

// Brute force recursion
func maxScore(points []int, k int) int {
    return maxScoreHelper(points, 0, len(points)-1, k)
}

func maxScoreHelper(points []int, i, j, k int) int {
    if k == 0 { return 0 }
    left := points[i] + maxScoreHelper(points, i+1, j, k-1)
    right := points[j] + maxScoreHelper(points, i, j-1, k-1)
    return max(left, right)
}

func max(x, y int) int { if x > y { return x }; return y }
