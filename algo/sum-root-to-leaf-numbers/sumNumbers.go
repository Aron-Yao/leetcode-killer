/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sum(root *TreeNode, now int, ret *int) {
    if root == nil {
        return
    }
    if root.Left == nil && root.Right == nil {
        *ret += now * 10 + root.Val
        return
    } 

    sum(root.Left, now * 10 + root.Val, ret)
    sum(root.Right, now * 10 + root.Val, ret)
}
 
func sumNumbers(root *TreeNode) int {
    ret := 0
    sum(root, 0, &ret)
    return ret
}
