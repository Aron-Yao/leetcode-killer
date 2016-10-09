/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    l := []*TreeNode{root}
    sum := 0
    
    for len(l) > 0 {
        x := l[0]
        l = l[1:]
        
        if x.Left != nil {
            if x.Left.Left == nil && x.Left.Right == nil {
                sum += x.Left.Val
            } else {
                l = append(l, x.Left)
            }
        }
        
        if x.Right != nil {
            if x.Right.Left != nil || x.Right.Right != nil {
                l = append(l, x.Right)
            }
        }
    }
    
    return sum
    
}
