/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    slow, fast := head, head
    var pre *ListNode
    
    for n > 0 && fast != nil {
        fast = fast.Next
        n--
    }
    
    for fast != nil {
        pre = slow
        slow, fast = slow.Next, fast.Next
    }
    if slow == head {
        return slow.Next
    }
    pre.Next = slow.Next
    return head
}
