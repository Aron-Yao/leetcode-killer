#define MAX(a, b) ((a) > (b) ? (a) : (b))

int integerBreak(int n) {
    int dp[(n > 4) ? (n + 1) : (5)];
    int i;
    
    dp[0] = 0, dp[1] = 1, dp[2] = 1, dp[3] = 2, dp[4] = 4;
    for (i = 5; i <= n; ++i) {
        dp[i] = 3 * MAX(i - 3, dp[i - 3]);
    }
    return dp[n];
}
