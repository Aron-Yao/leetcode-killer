func singleNumber(nums []int) []int {
    m := 0
    ret := []int{0, 0}
    
    for _, n := range nums {
        m ^= n
    }
    
    m &= -m
    
    for _, n := range nums {
        if (n & m) == 0 {
            ret[0] ^= n
        } else {
            ret[1] ^= n
        }
    }
    
    return ret
}
