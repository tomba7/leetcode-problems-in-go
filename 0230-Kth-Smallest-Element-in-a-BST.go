/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    var count int
    node := kthSmallestHelper(root, k, &count)
    return node.Val
}

func kthSmallestHelper(root *TreeNode, k int, count *int) *TreeNode {
    if root == nil { return nil }
    left := kthSmallestHelper(root.Left, k, count)
    if left != nil { return left }
    *count++
    if *count == k { return root }
    return kthSmallestHelper(root.Right, k, count)
}

---

func kthSmallest(root *TreeNode, k int) int {
    var res int
    kthSmallestHelper(root, &k, &res)
    return res
}

func kthSmallestHelper(root *TreeNode, k *int, res *int) bool {
    if root == nil { return false }
    if kthSmallestHelper(root.Left, k, res) {
        return true
    }
    *k--
    if *k == 0 {
        *res = root.Val
        return true
    }
    return kthSmallestHelper(root.Right, k, res)
}

---

func kthSmallest(root *TreeNode, k int) int {
    numLeftNodes := countNodes(root.Left)
    if k <= numLeftNodes {
        return kthSmallest(root.Left, k)
    } else if k == numLeftNodes + 1 {
        return root.Val
    }
    return kthSmallest(root.Right, k - numLeftNodes - 1)
}

func countNodes(node *TreeNode) int {
    if node == nil {return 0}
    return 1 + countNodes(node.Left) + countNodes(node.Right)
}
