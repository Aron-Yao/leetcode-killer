func bs(nums []int, s, e int, target int) int {
    if e < s {
        return -1
    }
    mid := (s + e) / 2
    if nums[mid] == target {
        return mid
    }
    if s == e {
        return -1
    }
    if nums[mid] < target {
        if nums[mid] <= nums[e] && target > nums[e] {
            e = mid - 1
        } else {
            s = mid + 1
        }
    } else {
        if nums[mid] >= nums[s] && target < nums[s] {
            s = mid + 1
        } else {
            e = mid - 1
        }
    }
    return bs(nums, s, e, target)
}


func search(nums []int, target int) int {
    return bs(nums, 0, len(nums) - 1, target)
}
