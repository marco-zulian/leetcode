/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
    var hashMap = make(map[int]int)

    postOrderTraversal(root, hashMap)

    var result []int
    var maxCount int

    for key, val := range hashMap {
        if val == maxCount {
            result = append(result, key)
        } else if val > maxCount {
            result = nil
            maxCount = val
            result = append(result, key)
        }
    }

    return result
}

func postOrderTraversal(root *TreeNode, hashMap map[int]int) {
    if root == nil { return }

    postOrderTraversal(root.Left, hashMap)
    postOrderTraversal(root.Right, hashMap)

    if root.Left != nil { root.Val += root.Left.Val }
    if root.Right != nil { root.Val += root.Right.Val }
    
    hashMap[root.Val] += 1
}