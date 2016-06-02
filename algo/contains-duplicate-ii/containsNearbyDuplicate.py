class Solution(object):
    def containsNearbyDuplicate(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: bool
        """
        hm = {}
        for i, n in enumerate(nums):
            if n in hm:
                if i - hm[n] > k:
                    hm[n] = i
                else:
                    return True
            else:
                hm[n] = i
        return False
