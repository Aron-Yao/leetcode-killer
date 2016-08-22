import "sort"

func threeSum(nums []int) [][]int {
    ans := [][]int{}
    sort.Ints(nums)
    
    var tmp, i int
    
    for i < len(nums) {
        e := nums[i]
        m, n := i + 1, len(nums) - 1
        
        for m < n {
            if nums[m] + nums[n] == 0 - e {
                ans = append(ans, []int{e, nums[m], nums[n]})
                
                tmp = m + 1
                for tmp < n && nums[tmp] == nums[m] {
                    tmp++
                }
                m = tmp
                tmp = n - 1
                for tmp > m && nums[tmp] == nums[n] {
                    tmp--
                }
                n = tmp
    
            } else if nums[m] + nums[n] < 0 - e {
                m++
            } else {
                n--
            }
        }
        
        tmp = i + 1
        for tmp < len(nums) && nums[tmp] == e {
            tmp++
        }
        i = tmp
    }
    return ans
}
