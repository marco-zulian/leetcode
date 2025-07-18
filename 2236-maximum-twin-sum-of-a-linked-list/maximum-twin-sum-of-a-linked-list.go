/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    fast, slow := head, head

    i := 0
    for fast != nil {
        fast = fast.Next.Next
        slow = slow.Next
        i += 1
    }

    sums := make([]int, i)
    j := 0
    for slow != nil {
        sums[j] += head.Val
        sums[i - 1 - j] += slow.Val

        slow = slow.Next
        head = head.Next
        j += 1
    }

    maxSum := 0
    for _, sum := range sums {
        if sum > maxSum {
            maxSum = sum
        }
    }

    return maxSum
}