class Solution(object):
    def containsNearbyAlmostDuplicate(self, nums, k, t):
        """
        :type nums: List[int]
        :type k: int
        :type t: int
        :rtype: bool
        """
        if t < 0:
            return False
            
        w = t + 1
        hm = {}
        for i, e in enumerate(nums):
            b = e / w
            if b in hm:
                return True
            if b - 1 in hm and abs(hm[b - 1] - e) <= t:
                return True
            if b + 1 in hm and abs(hm[b + 1] - e) <= t:
                return True
            hm[b] = e
            if i >= k:
                del hm[nums[i - k] / w]
        return False
