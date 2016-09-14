func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func rob(nums []int) int {
    last, now := 0, 0
    for _, n := range nums {
        last, now = now, max(now, last+n)
    }
    return now
}

