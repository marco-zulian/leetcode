/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    return checkSimetry(root.Left, root.Right)
}

func checkSimetry(left *TreeNode, right *TreeNode) bool {
    if left == nil && right == nil { return true }
    if left == nil || right == nil { return false }
    if left.Val != right.Val { return false }

    if !checkSimetry(left.Left, right.Right) { return false }
    return checkSimetry(left.Right, right.Left)
}