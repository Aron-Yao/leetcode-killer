func removeElement(nums []int, val int) int {
    p, q := 0, len(nums) - 1
    
    for p <= q {
        if nums[q] == val {
            q--
        } else {
            if nums[p] == val {
                nums[p], nums[q] = nums[q], nums[p]
                q--
            } 
            p++
        }
    }
    
    if q < 0 {
        return 0
    } else {
        return q+1
    }
}
