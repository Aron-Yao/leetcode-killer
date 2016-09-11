class Solution(object):
    def wordBreak(self, s, wordDict):
        """
        :type s: str
        :type wordDict: Set[str]
        :rtype: bool
        """ 
        table = [False] * (len(s) + 1)
        table[0] = True
        for i in xrange(1, len(s)+1):
            for j in xrange(i):
                if table[j] and s[j:i] in wordDict:
                    table[i] = True
                    break
        return table[-1]
