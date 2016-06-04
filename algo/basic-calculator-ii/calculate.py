class Solution(object):
    def doCalc(self, num1, num2, opr):
        if opr == '+':
            return num1 + num2
        elif opr == '-':
            return num1 - num2
        elif opr == '*':
            return num1 * num2
        else:
            return num1 / num2
    
    def calculate(self, s):
        """
        :type s: str
        :rtype: int
        """
        a, b = None, None
        opr1, opr2 = None, None
        n = 0
        length = len(s)
        while n < length:
            if s[n] == ' ':
                n += 1
            elif s[n] in ('+', '-', '*', '/'):
                if not opr1:
                    opr1 = s[n]
                else:
                    opr2 = s[n]
                n += 1
            elif s[n] >= '0' and s[n] <= '9':
                num = int(s[n])
                m = n + 1
                while m < length and s[m] >= '0' and s[m] <= '9':
                    num *= 10
                    num += int(s[m])
                    m += 1
                n = m
                if a == None:
                    a = num
                elif b == None:
                    b = num
                else:
                    if opr1 in ('+', '-') and opr2 in ('*' , '/'):
                        b = self.doCalc(b, num, opr2)
                        opr2 = None
                    else:
                        a, b = self.doCalc(a, b, opr1), num
                        opr1, opr2 = opr2, None
            else:
                n += 1
        return a if b == None else self.doCalc(a, b, opr1)
        
