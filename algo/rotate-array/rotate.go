func swap(nums []int, k, n int) {
    for i := 0; i < n; i++ {
        nums[i], nums[k+i] = nums[k+i], nums[i]
    }
}

func _rotate(nums []int, k int) {
    len1, len2 := len(nums[:k]), len(nums[k:]) 
    
    if len1 == 0 || len2 == 0 {
        return
    }
    if len1 == len2 {
        swap(nums, k, len1)
        return
    }
    if len1 > len2 {
        swap(nums, k, len2)
        _rotate(nums[len2:], len1 - len2)
    }
    if len2 > len1 {
        swap(nums, k, len1)
        _rotate(nums[len1:], len1)
    }
}

func rotate(nums []int, k int)  {
    _rotate(nums, len(nums) - (k % len(nums)))
}
