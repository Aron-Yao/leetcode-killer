# Definition for singly-linked list with a random pointer.
# class RandomListNode(object):
#     def __init__(self, x):
#         self.label = x
#         self.next = None
#         self.random = None

class Solution(object):
    def copyRandomList(self, head):
        """
        :type head: RandomListNode
        :rtype: RandomListNode
        """
        book = {}
        dummy = RandomListNode(0)
        tail = dummy
        p = head

        while p:
            new = RandomListNode(p.label)
            tail.next = new
            tail = tail.next
            book[p] = new
            p = p.next
        
        p = head
        
        while p:
            if p.random:
                book[p].random = book[p.random]
            p = p.next
            
        return dummy.next
        
        
