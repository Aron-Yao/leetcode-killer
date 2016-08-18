func getSum(a int, b int) int {
    if a&b != 0 {
        return getSum((a&b)<<1, a^b)
    } else {
        return a|b
    }
}
