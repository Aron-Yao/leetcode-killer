class Solution(object):
    def reverseWords(self, s):
        """
        :type s: str
        :rtype: str
        """
        str = ''
        
        for x in s.split():
            str += x[::-1] + ' '
            
        return str[:-1]
            
