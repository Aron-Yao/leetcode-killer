func bs(nums []int, target int) bool {
    start, end := 0, len(nums) - 1
    if target < nums[0] || target > nums[len(nums)-1] {
        return false
    }
    for start <= end {
        mid := (start + end) / 2
        if target == nums[mid] {
            return true
        }
        if target < nums[mid] {
            end = mid - 1
        }
        if target > nums[mid] {
            start = mid + 1
        }
    }
    return false
}

func search(nums []int, target int) bool {
    rp := len(nums)
    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] > nums[i+1] {
            rp = i + 1
            break
        }
    }
    if rp == len(nums) {
        return bs(nums, target)
    } else {
        return bs(nums[:rp], target) || bs(nums[rp:], target)
    }
}
