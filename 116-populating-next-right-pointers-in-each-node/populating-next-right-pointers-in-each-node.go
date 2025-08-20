/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    levelTraversal := []*Node{}
    
    var connectHelper func(*Node, int)
    connectHelper = func(root *Node, depth int) {
        if root == nil { return }
        
        if len(levelTraversal) == depth {
            levelTraversal = append(levelTraversal, root)
        } else {
            levelTraversal[depth].Next = root
            levelTraversal[depth] = levelTraversal[depth].Next
        }
        
        connectHelper(root.Left, depth+1)
        connectHelper(root.Right, depth+1)
    }
    
    connectHelper(root, 0)
    return root
}