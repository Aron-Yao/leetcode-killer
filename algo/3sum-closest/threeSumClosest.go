import "sort"

func abs(n int) int {
    if n >= 0 {
        return n
    } else {
        return -n
    }
}

func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    var i, j, k int
    var ret int = nums[0] + nums[1] + nums[2]
    
    for i = 0; i < len(nums) - 2; i++ {
        j, k = i + 1, len(nums) - 1
        for j < k {
            sum := nums[i] + nums[j] + nums[k]
            if sum == target {
                return sum
            }
            if abs(ret - target) > abs(sum - target) {
                ret = sum
            }
            if sum > target {
                k--
            }
            if sum < target {
                j++
            }
        }
    }
    
    return ret
}
