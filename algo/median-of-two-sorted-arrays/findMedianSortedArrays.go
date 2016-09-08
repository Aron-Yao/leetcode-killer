func min(a, b int) int {
    if a > b {
        return b
    } else {
        return a
    }
}

func find(nums1, nums2 []int, k int) int {
    if len(nums1) > len(nums2) {
        return find(nums2, nums1, k)
    }
    if len(nums1) == 0 {
        return nums2[k - 1]
    }
    if k == 1 {
        return min(nums1[0], nums2[0])
    }

    p := min(len(nums1), k / 2)
    q := k - p
    
    if nums1[p-1] < nums2[q-1] {
        return find(nums1[p:], nums2, k - p)
    } else if nums1[p-1] > nums2[q-1] {
        return find(nums1, nums2[q:], k - q)
    } else {
        return nums1[p-1]
    }
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    m, n := len(nums1), len(nums2)
    
    if (m + n) & 0x1 == 1 {
        return float64(find(nums1, nums2, (m + n) / 2 + 1))
    } else {
        return float64((find(nums1, nums2, (m + n) / 2) + find(nums1, nums2, (m + n) / 2 + 1))) / 2
    }
}
