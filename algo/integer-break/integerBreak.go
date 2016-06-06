func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func integerBreak(n int) int {
    dp := []int{0, 1, 1, 2, 4}
    for i := 5; i <= n; i++ {
        dp = append(dp, 3 * max(dp[i - 3], i - 3))
    }
    return dp[n]
}
