func sum(nums []int, start, k int) int {
    var ret int

    for _, n := range nums[start:start+k] {
        ret += n
    }
    
    return ret
}

func imax(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func findMaxAverage(nums []int, k int) float64 {
    var max int
    var tmp int
    var inited bool
    
    for i := 0; i + k <= len(nums); i++ {
        if !inited {
            max = sum(nums, i, k)
            inited = true
        } else {
            tmp = sum(nums, i, k)
            max = imax(max, tmp)
        }
    }
    
    return float64(max) / float64(k)
}
