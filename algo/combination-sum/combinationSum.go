import "sort"

func combine(candidates *[]int, target int, index int, now []int, ret *([][]int)) {
    for i := index; i < len(*candidates); i++ {
        n := (*candidates)[i]
        new_target := target - n
        if new_target == 0 {
            now = append(now, n)
            *ret = append(*ret, now)
            break
        } else if new_target > 0 {
            now = append([]int{}, now...)
            combine(candidates, new_target, i, append(now, n), ret)
        } else {
            break
        }
    }
}

func combinationSum(candidates []int, target int) [][]int {
    ret := [][]int{}
    sort.Ints(candidates)
    combine(&candidates, target, 0, []int{}, &ret)
    return ret
}