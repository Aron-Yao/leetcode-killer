/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    p := &ListNode{Val:0, Next:head}
    q := p
    
    for p.Next != nil {
        if p.Next.Val == val {
            p.Next = p.Next.Next
        } else {
            p = p.Next
        }
    }
    return q.Next
}
