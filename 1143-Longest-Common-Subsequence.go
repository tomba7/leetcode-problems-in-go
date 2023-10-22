/*
 s1 = "abcde"
          i
 s2 = "ace"
         j
            0,0
            1,1
          /      \
        2,1      1,2
         |
        3,2
       /   \ 
     4,2   3,3
      |
     5,3
 */
func longestCommonSubsequence(s1 string, s2 string) int {
    m, n := len(s1), len(s2)
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            if s1[i] == s2[j] {
                dp[i][j] = dp[i+1][j+1] + 1
            } else {
                dp[i][j] = max(dp[i+1][j], dp[i][j+1])
            }
        }
    }
    return dp[0][0]
}

func max(x, y int) int { if x > y { return x }; return y }

---

func longestCommonSubsequence(s1 string, s2 string) int {
    m, n := len(s1), len(s2)
    memo := make([][]int, m)
    for i := range memo {
        memo[i] = make([]int, n)
        for j := range memo[i] { memo[i][j] = -1 }
    }
    return longestCommonSubseqHelper(s1, s2, 0, 0, memo)
}

func longestCommonSubseqHelper(s1, s2 string, i, j int, memo [][]int) int {
    if i == len(s1) || j == len(s2) {
        return 0
    }
    if memo[i][j] != -1 {
        return memo[i][j]
    }
    if s1[i] == s2[j] {
        memo[i][j] = 1 + longestCommonSubseqHelper(s1, s2, i+1, j+1, memo)
    } else {
        left := longestCommonSubseqHelper(s1, s2, i+1, j, memo)
        right := longestCommonSubseqHelper(s1, s2, i, j+1, memo)
        memo[i][j] = max(left, right)
    }
    return memo[i][j]
}

func max(x, y int) int { if x > y { return x }; return y }
