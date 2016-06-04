func containsDuplicate(nums []int) bool {
    hm := map[int]bool{}
    for _, n := range nums {
        if _, ok := hm[n]; ok {
            return true
        } else {
            hm[n] = true
        }
    }
    return false
}
