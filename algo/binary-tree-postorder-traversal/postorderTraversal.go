/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    stack, ret := []*TreeNode{root}, []int{}
    var now, pre *TreeNode
    
    for len(stack) > 0 {
        now = stack[len(stack)-1]
        if (now.Left == nil && now.Right == nil) || (pre != nil && (pre == now.Left || pre == now.Right)) {
            pre = now
            ret = append(ret, now.Val)
            stack = stack[:len(stack)-1]
        } else {
            if now.Right != nil {
                stack = append(stack, now.Right)
            }
            if now.Left != nil {
                stack = append(stack, now.Left)
            }
        }
    }
    
    return ret
}
