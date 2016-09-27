func permute(nums []int) [][]int {
    ret := [][]int{[]int{nums[0]}}
    
    for _, n := range nums[1:] {
        new_ret := make([][]int, 0, len(ret)*(len(ret[0])+1))
        for _, m := range ret {
            for i := 0; i <= len(m); i++ {
                tmp := make([]int, 0, len(m)+1)
                tmp = append(tmp, m[:i]...)
                tmp = append(tmp, n)
                tmp = append(tmp, m[i:]...)
                new_ret = append(new_ret, tmp)
            }
        }
        ret = new_ret
    }
    
    return ret
}
