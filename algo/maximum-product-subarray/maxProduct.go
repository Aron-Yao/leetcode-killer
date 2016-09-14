func Max(a, b int) int {
    if a >  b {
        return a
    } else {
        return b
    }
}

func Min(a, b int) int {
    if a > b {
        return b
    } else {
        return a
    }
}

func swap(a, b *int) {
    *a, *b = *b, *a
}

func maxProduct(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    ret := nums[0]
    max, min := ret, ret
        
    for _, x := range nums[1:] {
        if x < 0 {
            swap(&max, &min)
        }
        
        max = Max(x, x * max)
        min = Min(x, x * min)
        ret = Max(ret, max)
    }    
    return ret
}
