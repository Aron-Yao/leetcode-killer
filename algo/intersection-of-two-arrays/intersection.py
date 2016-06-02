class Solution(object):
    def intersection(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: List[int]
        """
        hm, ret = {}, []
        for n in nums1:
            if n not in hm:
                hm[n] = 1
        for n in nums2:
            if n in hm:
                del hm[n]
                ret.append(n)
        return ret
                