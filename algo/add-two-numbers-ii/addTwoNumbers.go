/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDiffBits(l1, l2 *ListNode) int {
    p, q := l1, l2
    ret := 0
    
    for p != nil && q != nil {
        p, q = p.Next, q.Next 
    }
    
    if p == nil && q == nil {
        return ret
    }
    
    if p == nil {
        for q != nil {
            q = q.Next
            ret--
        }
    } else {
        for p != nil {
            p = p.Next
            ret++
        }
    }
    
    return ret
}

func getStartPoint(l1, l2 *ListNode, diff int) (l1n, l2n *ListNode) {
    l1n, l2n = l1, l2
    
    if diff == 0 {
        return
    }
    
    if diff > 0 {
        for diff > 0 {
            l1n = l1n.Next
            diff--
        }
    } else {
        for diff < 0 {
            l2n = l2n.Next
            diff++
        }
    }
    
    return
}

func doSum(l1, l2 *ListNode, l1n, l2n *ListNode) *ListNode {
    head := &ListNode{}
    tail := head    
    p, q := l1, l2
    carry := false
    
    var prev, next *ListNode
    
    if l1 == l1n {
        for q != l2n {
            tail.Next = &ListNode{Val: q.Val}
            tail = tail.Next
            q = q.Next
        }
    }
    
    if l2 == l2n {
        for p != l1n {
            tail.Next = &ListNode{Val: p.Val}
            tail = tail.Next
            p = p.Next
        }
    }
    
    for p != nil && q != nil {
        tail.Next = &ListNode{Val: p.Val+q.Val}
        tail = tail.Next
        p, q = p.Next, q.Next
    }
    
    prev = head
    next = prev.Next
    
    for carry || next != nil {
        if next == nil && carry {
            carry = false
            prev = head
            next = prev.Next
        }
        
        if next.Val >= 10 {
            if prev == head {
                head.Next = &ListNode{Val: 1, Next: next}
                prev = head.Next
            } else {
                prev.Val += 1
                carry = true
            }
            
            next.Val -= 10
        }
        
        prev, next = prev.Next, next.Next
    }
        
    return head.Next
}
            
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    diff := getDiffBits(l1, l2)
    l1n, l2n := getStartPoint(l1, l2, diff)
    
    return doSum(l1, l2, l1n, l2n)
}
