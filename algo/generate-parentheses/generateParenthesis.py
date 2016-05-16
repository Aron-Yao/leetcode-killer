class Solution(object):
    def add(self, l, s, n, m):
        if n == 0 and m == 0:
            l.append(s)
        else:
            if m > 0: self.add(l, s + ')', n, m - 1)
            if n > 0: self.add(l, s + '(', n - 1, m + 1)
    
    def generateParenthesis(self, n):
        """
        :type n: int
        :rtype: List[str]
        """
        l = []
        self.add(l, '', n, 0)
        return l