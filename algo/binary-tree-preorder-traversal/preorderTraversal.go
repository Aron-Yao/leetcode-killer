/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    stack, ret := []*TreeNode{}, []int{}
    p := root
    
    for p != nil || len(stack) > 0 {
        for p != nil {
            ret = append(ret, p.Val)
            stack = append(stack, p)
            p = p.Left
        }
        
        if len(stack) > 0 {
            p = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            p = p.Right
        }
    }
    
    return ret
    
}
