func Max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func maxSubArray(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    dp := make([]int, len(nums))
    max := nums[0]
    dp[0] = nums[0]
    
    for i := 1; i < len(nums); i++ {
        n := nums[i]
        if dp[i-1] > 0 {
            dp[i] = dp[i-1] + n
        } else {
            dp[i] = n
        }
        max = Max(max, dp[i])
    }
    return max
}
