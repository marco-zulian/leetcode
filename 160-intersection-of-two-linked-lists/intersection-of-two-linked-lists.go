/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    visited := map[*ListNode]bool{}

    i := headA
    for i != nil {
        visited[i] = true
        i = i.Next
    }

    j := headB
    for j != nil {
        if _, ok := visited[j]; ok { return j }
        j = j.Next
    }

    return nil
}