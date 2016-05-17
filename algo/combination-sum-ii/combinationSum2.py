class Solution(object):
    def combine(self, candidates, index, target, now, ret):
        for x in xrange(index, len(candidates)):
            n = candidates[x]
            if n > target:
                return
            if n == target:
                new = now + [n]
                if new not in ret:
                    ret.append(now + [n])
                return
            self.combine(candidates, x + 1, target - n, now + [n], ret)
    
    def combinationSum2(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        candidates.sort()
        ret = []
        self.combine(candidates, 0, target, [], ret)
        return ret