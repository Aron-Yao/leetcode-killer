func intersection(nums1 []int, nums2 []int) []int {
    hm := map[int]int{}
    ret := []int{}
    for _, n := range nums1 {
        if _, ok := hm[n]; !ok {
            hm[n] = 1
        }
    }
    for _, n := range nums2 {
        if _, ok := hm[n]; ok {
            delete(hm, n)
            ret = append(ret, n)
        }
    }
    return ret
}