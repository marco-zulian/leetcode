/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    fast, slow := head, head

    for {
        fast = fast.Next
        if fast == nil {
            return slow
        }
        slow = slow.Next
        fast = fast.Next
        if fast == nil {
            return slow
        }
    }

    return nil
}