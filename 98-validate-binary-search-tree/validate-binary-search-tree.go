/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    _, valid := inOrderTraverseBSTValid(root, int(math.Inf(-1)))
    return valid
}

func inOrderTraverseBSTValid(root *TreeNode, curr int) (int, bool) {
    if root == nil { return curr, true }
    
    curr, valid := inOrderTraverseBSTValid(root.Left, curr)
    
    if !valid || root.Val <= curr { return curr, false }
    curr = root.Val

    return inOrderTraverseBSTValid(root.Right, curr)
}