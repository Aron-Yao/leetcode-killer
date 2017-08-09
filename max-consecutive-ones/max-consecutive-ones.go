func findMaxConsecutiveOnes(nums []int) int {
    var max, cnt int
    
    for _, n := range nums {
        if n == 1 {
            cnt++
        } else {
            if cnt > max {
                max = cnt
            }
            cnt = 0
        }
    }
    
    if cnt > max { 
        max = cnt
    }
    
    return max
}
