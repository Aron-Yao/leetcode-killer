func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func findMin(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    if len(nums) == 2 {
        return min(nums[0], nums[1])
    }
    
    s, e := 0, len(nums) - 1
    mid := (s + e) / 2
    
    if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
        return nums[mid+1]
    }
    if nums[mid] < nums[mid-1] && nums[mid] < nums[mid+1] {
        return nums[mid]
    }
    if nums[mid] < nums[len(nums)-1] {
        return findMin(nums[:mid])
    } else {
        return findMin(nums[mid+1:])
    }
    return -1
}
