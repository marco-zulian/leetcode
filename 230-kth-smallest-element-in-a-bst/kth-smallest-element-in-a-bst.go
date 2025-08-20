/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    var helper func(*TreeNode, int, int) (int, int)
    helper = func(root *TreeNode, k int, ith int) (int, int) {
        if root == nil { return ith, -1 }
        
        ith, val := helper(root.Left, k, ith)
        if val != -1 { return ith, val }
        ith++
        if ith == k { return ith, root.Val }
        return helper(root.Right, k, ith)
    }
    
    _, val := helper(root, k, 0)
    return val
}