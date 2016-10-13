/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }
    
    var cnt int
    var p *ListNode = head
    var q *ListNode = nil
    
    for p != nil {
        cnt++
        q = p
        p = p.Next
    }
    
    q.Next = head
    
    k = cnt - k % cnt - 1
    p = head
    
    for k > 0 {
        p = p.Next
        k--
    }
    
    q = p.Next
    p.Next = nil
    
    return q
}
