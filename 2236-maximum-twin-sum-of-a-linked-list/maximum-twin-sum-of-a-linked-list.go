/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    fast, slow := head, head
    sums := []int{}

    i := 0
    for fast != nil {
        fast = fast.Next.Next
        sums = append(sums, slow.Val)
        slow = slow.Next
        i += 1
    }

    j, maxSum := 0, 0
    for slow != nil {
        sums[i - 1 - j] += slow.Val
        if sums[i - 1 - j] > maxSum {
            maxSum = sums[i - 1 -j]
        }

        j += 1
        slow = slow.Next
    }

    return maxSum
}