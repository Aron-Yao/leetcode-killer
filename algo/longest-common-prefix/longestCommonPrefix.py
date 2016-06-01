class Solution(object):
    def longestCommonPrefix(self, strs):
        """
        :type strs: List[str]
        :rtype: str
        """
        if not strs:
            return ""
        
        last, lengths = 0, {}
        for x in xrange(1, len(strs)):
            lengths[x] = len(strs[x])
            
        for x in xrange(len(strs[0])):
            saved = strs[0][x]
            for y in xrange(1, len(strs)):
                if x >= lengths[y] or saved != strs[y][x]:
                    return strs[0][:last]
            last += 1
        return strs[0]