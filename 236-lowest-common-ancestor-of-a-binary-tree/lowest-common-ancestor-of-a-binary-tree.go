import (
    "fmt"
    "maps"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    fmt.Println(p.Val, q.Val)
    _, desc := traverse(root, p, q)
    return desc
}

func traverse(root, p, q *TreeNode) (map[int]bool, *TreeNode) {
    if (root == nil) { return map[int]bool{}, nil }

    decendants := map[int]bool{}

    leftDesc, result := traverse(root.Left, p, q)
    if result != nil { return map[int]bool{}, result }
    maps.Copy(decendants, leftDesc)

    rightDesc, result := traverse(root.Right, p, q)
    if result != nil { return map[int]bool{}, result }
    maps.Copy(decendants, rightDesc)

    decendants[root.Val] = true
    _, isP := decendants[p.Val]
    _, isQ := decendants[q.Val]

    if isP && isQ { return map[int]bool{}, root }
    return decendants, nil
}