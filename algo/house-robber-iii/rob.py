# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def doRob(self, root):
        if not root:
            return [0, 0]
        left, right = self.doRob(root.left), self.doRob(root.right)
        return [max(left[0], left[1]) + max(right[0], right[1]), root.val + left[0] + right[0]]
    
    
    def rob(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        return max(self.doRob(root))
        
