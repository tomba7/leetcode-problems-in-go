// Given a binary tree, check whether it is a mirror of itself (ie,
// symmetric around its center).
//
// For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
//
//     1
//    / \
//   2   2
//  / \ / \
// 3  4 4  3
//
// But the following [1,2,2,null,3,null,3] is not:
//
//     1
//    / \
//   2   2
//    \   \
//    3    3
//
// Note:
// Bonus points if you could solve it both recursively and iteratively.

func isSymmetric(root *TreeNode) bool {
    return root == nil || isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(left, right *TreeNode) bool {
    if left == nil && right == nil {
        return true
    } else if left == nil || right == nil {
        return false
    }
    return left.Val == right.Val &&
            isSymmetricHelper(left.Left, right.Right) &&
            isSymmetricHelper(left.Right, right.Left)
}
