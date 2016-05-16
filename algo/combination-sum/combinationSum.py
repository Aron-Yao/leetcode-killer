class Solution(object):
    def combine(self, target, index, now, ret):
        for x in xrange(index, self.length):
            n = self.candidates[x]
            new_target = target - n
            if new_target == 0:
                ret.append(now + [n])
                return
            elif new_target < 0:
                return
            else:
                self.combine(new_target, x, now + [n], ret)
        
    def combinationSum(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        self.candidates = sorted(candidates)
        self.length = len(self.candidates)
        ret = []
        self.combine(target, 0, [], ret)
        return ret
        