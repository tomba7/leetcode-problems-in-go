/*
                        6              
                      /   \
                    /       \
                  2           8
                /   \       /   \
               0     4     7     9
                   /   \
                  3     5
              
              p = 2
              q = 4
 
 Algorithm
 * If p.Val > q.Val: swap(p, q): to make sure p is the left subtree
 * User iteration starting from the root
 * At every iteration:
   - If p.Val < curr.Val && q.Val < curr.Val:
        curr = curr.Left
   - If p.Val > curr.Val && q.Val > curr.Val
        curr = curr.Right
   - else
        lca = curr
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    curr := root
    if p.Val > q.Val { p, q = q, p }
    for curr != nil {
        if q.Val < curr.Val {
            curr = curr.Left
        } else if p.Val > curr.Val {
            curr = curr.Right
        } else {
            return curr
        }
    }
    return nil
}
