class Token_Kind(object):
    BLANK, SIGN, NUM, INT, EPX = xrange(5)

class Solution(object):
    def eat(self, s, start, kind):
        if len(s[start:]) == 0:
            return start
        
        if kind == Token_Kind.BLANK:
            while start < len(s) and s[start] == ' ':
                start += 1
            return start
        elif kind == Token_Kind.SIGN:
            return start + 1 if s[start] == '+' or s[start] == '-' else start
        elif kind == Token_Kind.NUM:
            nstart = self.eat(s, start, Token_Kind.INT)
            hasNum = True if nstart != start else False
            if nstart < len(s) and s[nstart] == '.':
                nstart += 1
            nnstart = self.eat(s, nstart, Token_Kind.INT)
            return start if not hasNum and nstart == nnstart else nnstart
        elif kind == Token_Kind.INT:
            while start < len(s) and s[start] >= '0' and s[start] <= '9':
                start += 1
            return start
        elif kind == Token_Kind.EPX:
            return start + 1 if s[start] == 'e' else start
        else:
            return start
        
    def isNumber(self, s):
        """
        :type s: str
        :rtype: bool
        """
        if not s:
            return false
        
        old = self.eat(s, self.eat(s, 0, Token_Kind.BLANK), Token_Kind.SIGN)
        new = self.eat(s, old, Token_Kind.NUM)
        if new != old:
            old, new = new, self.eat(s, new, Token_Kind.EPX)
            if new != old:
                new = self.eat(s, new, Token_Kind.SIGN)
                old, new = new, self.eat(s, new, Token_Kind.INT)
                if old == new:
                    return False
            new = self.eat(s, new, Token_Kind.BLANK)
            if new == len(s):
                return True 
        return False
        