func doRemove(nums []int, n int) int {
    if len(nums) <= n {
        return len(nums)
    }
    
    var p, q, cnt, now int
    
    now = nums[0]
    
    for q < len(nums) {
        if nums[q] == now {
            cnt++
        } else {
            now = nums[q]
            cnt = 1
        }
        if cnt <= n {
            nums[p] = now
            p++
        }
        q++
    }
    
    return p
}

func removeDuplicates(nums []int) int {
    return doRemove(nums, 2)
}
