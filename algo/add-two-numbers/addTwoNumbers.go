/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    tmp := 0
    res := new(ListNode)
    tail := res
    
    for l1 != nil || l2 != nil {
        tmp /= 10
        if l1 != nil {
            tmp += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            tmp += l2.Val
            l2 = l2.Next
        }
        tail.Next = &ListNode{tmp % 10, nil}
        tail = tail.Next
    }
    if tmp / 10 == 1 {
        tail.Next = &ListNode{1, nil}
    }
    return res.Next
}
