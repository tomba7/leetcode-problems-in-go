/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) (res []string) {
    path := []byte{}
    binaryTreePathsHelper(root, path, &res)
    return
}

func binaryTreePathsHelper(node *TreeNode, path []byte, res *[]string) {
    if node == nil { return }
    if node.Left == nil && node.Right == nil {
        path = append(path, fmt.Sprintf("%d", node.Val)...)
        *res = append(*res, string(path))
        return
    }
    path = append(path, fmt.Sprintf("%d->", node.Val)...)
    binaryTreePathsHelper(node.Left, path, res)
    binaryTreePathsHelper(node.Right, path, res)
}

---

func binaryTreePaths(root *TreeNode) []string {
    var res []string
    binaryTreePathsHelper(root, nil, &res)
    return res
}

func binaryTreePathsHelper(root *TreeNode, s []byte, res *[]string) {
    if root == nil { return }
    s = append(s, fmt.Sprintf("%d", root.Val)...)
    if root.Left == nil && root.Right == nil {
        *res = append(*res, string(s))
    } else {
        s = append(s, fmt.Sprintf("->")...)
        binaryTreePathsHelper(root.Left, s, res)
        binaryTreePathsHelper(root.Right, s, res)
    }
}
