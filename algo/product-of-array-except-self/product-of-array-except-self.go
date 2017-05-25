func productExceptSelf(nums []int) []int {
    all, zcnt, zi := 1, 0, -1
    ret := make([]int, len(nums))
    
    for i, n := range nums {
        if n == 0 {
            zcnt++
            if zcnt == 1 {
                zi = i
            }
            continue
        }
        all *= n
    }
    
    if zcnt >=2 {
        // all 0
    } else if zcnt == 1 {
        ret[zi] = all
    } else {
        for i, n := range nums {
            ret[i] = all / n
        }
    }
    
    return ret
}
