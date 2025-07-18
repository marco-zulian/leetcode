/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortLinkedList(head *ListNode) *ListNode {
    i, j := head, head

    for j.Next != nil {
        if j.Next.Val < 0 {
            k := j.Next
            j.Next = j.Next.Next
            k.Next = i
            i = k
        } else {
            j = j.Next
        }
    }

    return i
}