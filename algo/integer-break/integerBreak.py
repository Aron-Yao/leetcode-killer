class Solution(object):
    def integerBreak(self, n):
        """
        :type n: int
        :rtype: int
        """
        dp = [0, 1, 1, 2, 4]
        for x in range(5, n + 1):
            dp.append(3 * max(dp[x - 3], x - 3))
        return dp[n] 
        
