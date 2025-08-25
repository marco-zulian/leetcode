/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    dummy := &ListNode{
        Val: 0,
        Next: head,
    }
    curr, greatest := head.Next, head
    
    for curr != nil {
        insertPos, prev := dummy.Next, dummy
        for insertPos != nil && insertPos != curr && insertPos.Val < curr.Val {
            insertPos = insertPos.Next
            prev = prev.Next
        }
        
        aux := curr.Next
        curr.Next = insertPos
        prev.Next = curr
        if curr.Val > greatest.Val {
            greatest = curr
        }
        greatest.Next = aux
        curr = aux
    }
    
    return dummy.Next
}