/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    result := [][]int{}
    return inOrderTraversal(root, result, 0)
}

func inOrderTraversal(root *TreeNode, currList [][]int, depth int) [][]int {
    if root == nil { return currList }

    if len(currList) == depth {
        currList = append(currList, []int{root.Val})
    } else {
        currList[depth] = append(currList[depth], root.Val)
    }

    currList = inOrderTraversal(root.Left, currList, depth+1)
    currList = inOrderTraversal(root.Right, currList, depth+1)

    return currList
}