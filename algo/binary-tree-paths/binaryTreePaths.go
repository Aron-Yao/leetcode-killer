/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func dfs(root *TreeNode, path string, res *[]string) {
    if root == nil {
        // do nothing
        return
    }
    if root.Left == nil && root.Right == nil {
        *res = append(*res, path + strconv.Itoa(root.Val))
    }
    if root.Left != nil {
        dfs(root.Left, path + strconv.Itoa(root.Val) + "->", res)
    }
    if root.Right != nil {
        dfs(root.Right, path + strconv.Itoa(root.Val) + "->", res)
    }
} 
 
func binaryTreePaths(root *TreeNode) []string {
    res := new([]string)
    dfs(root, "", res)
    return *res
}
