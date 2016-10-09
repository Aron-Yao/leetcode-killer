/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    stack, ret := []*TreeNode{}, []int{}
    var p *TreeNode = root
    
    for p != nil || len(stack) > 0 {
        for p != nil {
            stack = append(stack, p)
            p = p.Left
        }
        
        if len(stack) > 0 {
            p = stack[len(stack)-1]
            ret = append(ret, p.Val)
            p = p.Right
            stack = stack[:len(stack)-1]
        }
    }

    return ret
    
}
