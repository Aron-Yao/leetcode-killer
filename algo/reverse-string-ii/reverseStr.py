class Solution(object):
    def reverseStr(self, s, k):
        """
        :type s: str
        :type k: int
        :rtype: str
        """
        retStr = ''
        start = 0
        
        while start < len(s):
            retStr += s[start:start+k][::-1] + s[start+k:start+2*k]
            start += 2*k
            
        return retStr
