func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0
    }
    
    var total, tend, buy int
    
    buy = -1
    
    for i := 1; i < len(prices); i++ {
        if prices[i] > prices[i-1] && tend != 1 {
            tend = 1
            if buy == -1 {
                buy = prices[i-1]
            }
        }
        if prices[i] == prices[i-1] {
            tend = 0
        }
        if prices[i] < prices[i-1] && tend != -1 {
            tend = -1
            if buy != -1 {
                total += prices[i-1] - buy
                buy = -1
            }
        }
    }
    if buy != -1 && buy < prices[len(prices)-1] {
        total += prices[len(prices)-1] - buy
    }
    return total
}
