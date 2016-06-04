class Solution(object):
    def containsDuplicate(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        hm = {}
        for n in nums:
            if n in hm:
                return True
            else:
                hm[n] = 1
        return False
