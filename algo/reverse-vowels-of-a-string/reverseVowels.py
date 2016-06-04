class Solution(object):
    def isVowels(self, c):
        return c in ('a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U')
        
    def reverseVowels(self, s):
        """
        :type s: str
        :rtype: str
        """
        if s == None:
            return None
            
        l = list(s)
        start, end = 0, len(s) - 1
        while start < end:
            if self.isVowels(l[start]):
                while start < end and not self.isVowels(l[end]):
                    end -= 1
                l[start], l[end] = l[end], l[start]
                end -= 1
            start += 1
        return "".join(l)
        
        
