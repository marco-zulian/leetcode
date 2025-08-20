/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    i, j := l1, l2
    
    for i.Next != nil && j.Next != nil {
        val := (i.Val + j.Val + carry) % 10
        carry = (i.Val + j.Val + carry) / 10
        
        i.Val = val
        i = i.Next
        j = j.Next
    }
    
    val := (i.Val + j.Val + carry) % 10
    carry = (i.Val + j.Val + carry) / 10
        
    i.Val = val
    
    if i.Next == nil && j.Next != nil {
        j = j.Next
        i.Next = j
    }
    
    if i.Next != nil {
        i = i.Next

        for i.Next != nil {
            val := (i.Val + carry) % 10
            carry = (i.Val + carry) / 10

            i.Val = val
            i = i.Next
        }

        val = (i.Val + carry) % 10
        carry = (i.Val + carry) / 10

        i.Val = val
    }

    if carry > 0 {
        i.Next = &ListNode{Val: 1, Next: nil}
    }

    return l1
}