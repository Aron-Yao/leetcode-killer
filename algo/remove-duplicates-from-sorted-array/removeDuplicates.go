func removeDuplicates(nums []int) int {
    if len(nums) < 2 {
        return len(nums)
    }
    
    p, q := 0, 1
    for p < len(nums) && q < len(nums) {
        if nums[p] == nums[q] {
            q++
        } else {
            p++
            nums[p] = nums[q]
            q++
        }
    }
    return p+1
}
