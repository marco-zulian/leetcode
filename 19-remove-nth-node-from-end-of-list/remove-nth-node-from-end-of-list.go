/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    i, j := head, head

    for _ = range n {
        j = j.Next
    }

    if j == nil {
        return head.Next
    }

    for j.Next != nil {
        j = j.Next
        i = i.Next
    }

    if i.Next.Next == nil {
        i.Next = nil
    } else {
        i.Next = i.Next.Next
    }

    return head
}