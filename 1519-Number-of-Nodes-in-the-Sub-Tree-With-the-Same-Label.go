/*
 - Create a directed graph from edges (no need for seen)
 - Starting at every index from 0 -> n-1, perform DFS and count
   the number of matching labels in its subtree
   
   n = 4, edges = [[0,1],[1,2],[0,3]], labels = "bbbb"
   0: [1, 3]
   1: [2]
   
   [[0,2],[0,3],[1,2]]
   0: [2, 3]    a
   1: [2]       e
   2: []        e
   3: []        d
                     0 a
                   /   \
                 2 e     3 d
                /
              1 e
 */
func countSubTrees(n int, edges [][]int, labels string) []int {
    graph := make([][]int, n)
    for _, edge := range edges {
        a, b := edge[0], edge[1]
        graph[a] = append(graph[a], b)
        graph[b] = append(graph[b], a)
    }
    res := make([]int, n)
    countLabels(0, labels, graph, make([]bool, n), res)
    return res
}

func countLabels(node int, labels string, graph [][]int, seen []bool, res []int) []int {
    count := make([]int, 26)
    if !seen[node] {
        seen[node] = true
        label := labels[node]
        count[label-'a']++
        for _, child := range graph[node] {
            childCount := countLabels(child, labels, graph, seen, res)
            for i := 0; i < len(childCount); i++ {
                count[i] += childCount[i]
            }
        }
        res[node] = count[label-'a']
    }
    return count
}
