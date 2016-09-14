func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func doRob(nums []int) int {
    last, now := 0, 0
    for _, n := range nums {
        last, now = now, max(now, last+n)
    }
    return now
}


func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    if len(nums) == 2 {
        return max(nums[0], nums[1])
    }
    return max(doRob(nums[:len(nums)-1]), doRob(nums[1:]))
}
