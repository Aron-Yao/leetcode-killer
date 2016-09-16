func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0
    }
    min := prices[0]
    ret := 0
    for i := 1; i < len(prices); i++ {
        if prices[i] - min > ret {
            ret = prices[i] - min
        }
        if prices[i] < min {
            min = prices[i]
        }
    }
    return ret
}
