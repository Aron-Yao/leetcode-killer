func intersect(nums1 []int, nums2 []int) []int {
    hm := map[int]int{}
    ret := []int{}
    for _, n := range nums1 {
        if _, ok := hm[n]; ok {
            hm[n] += 1
        } else {
            hm[n] = 1
        }
    }
    for _, n := range nums2 {
        if times, ok := hm[n]; ok && times > 0 {
            hm[n] -= 1
            ret = append(ret, n)
        }
    }
    return ret
}