func missingNumber(nums []int) int {
    res := 0
    for i, n := range nums {
        res += i + 1 - n
    }
    return res
}
