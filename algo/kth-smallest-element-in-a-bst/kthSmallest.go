/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    stack := []*TreeNode{}
    p := root
    cnt := 0
    
    for p != nil || len(stack) > 0 {
        for p != nil {
            stack = append(stack, p)
            p = p.Left
        }
        
        cnt++
        if cnt == k {
            return stack[len(stack)-1].Val
        } else {
            p = stack[len(stack)-1].Right
            stack = stack[:len(stack)-1]
        }
    }
    
    return 0
}
