/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    ret := []int{root.Val}
    cur, next := []*TreeNode{root}, []*TreeNode{}
    
    for len(cur) > 0 || len(next) > 0 {
        for len(cur) > 0 {
            p := cur[0]
            if p.Left != nil {
                next = append(next, p.Left)
            }
            if p.Right != nil {
                next = append(next, p.Right)
            }
            cur = cur[1:]
        }
        if len(next) > 0 {
            ret = append(ret, next[len(next)-1].Val)
        }
        cur, next = next, cur
    }
    
    return ret
}
