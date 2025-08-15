/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil { return false }

    i, j := head, head.Next 
    for j != nil {
        j = j.Next
        
        if j == i {
            return true
        }

        if j == nil {
            break
        }

        i = i.Next
        j = j.Next
    }

    return false
}