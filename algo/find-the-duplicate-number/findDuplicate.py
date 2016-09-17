class Solution(object):
    def findDuplicate(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        slow, fast, meet = nums[0], nums[nums[0]], 0
        
        while slow != fast:
            slow, fast = nums[slow], nums[nums[fast]]
        
        while meet != slow:
            slow, meet = nums[slow], nums[meet]
            
        return meet
