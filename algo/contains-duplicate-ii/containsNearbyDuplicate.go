func containsNearbyDuplicate(nums []int, k int) bool {
    hm := map[int]int{}
    for i, n := range nums {
        if e, ok := hm[n]; ok {
            if i - e > k {
                hm[n] = i
            } else {
                return true
            }
        } else {
            hm[n] = i
        }
    }
    return false
}
