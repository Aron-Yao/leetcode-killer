func hammingDistance(x int, y int) int {
    cnt := 0
    m := x ^ y
    
    for m > 0 {
        cnt += m & 0x1
        m = m >> 1
    }
    
    return cnt
}
