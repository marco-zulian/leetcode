/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head.Next == nil { return true }
    middleNode := findMiddle(head)
    reversedHead := reverse(middleNode)
    isPalindrome := startsWith(head, reversedHead)    
    
    return isPalindrome
}

func findMiddle(head *ListNode) *ListNode {
    i, j := head, head

    for j != nil {
        j = j.Next
        i = i.Next

        if j == nil { break }
        j = j.Next
    }

    return i
}

func reverse(head *ListNode) *ListNode {
    var prev *ListNode

    prev, curr, next := nil, head, head.Next

    for next != nil {
        curr.Next = prev
        prev, curr, next = curr, next, next.Next
    }
    curr.Next = prev

    return curr
}

func startsWith(head1 *ListNode, head2 *ListNode) bool {
    i, j := head1, head2
    for j != nil {
        if i.Val != j.Val { return false }
        j = j.Next
        i = i.Next
    }

    return true
}