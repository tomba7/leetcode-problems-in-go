/*
n = 4
0: [1, 2]
1: [3]
2: [3]
3: []

- Create a graph from prerequisites
- Perform DFS from every node (0-numCourses-1)
- If a cycle is detected (using node coloring) then return false
- If not return true
*/

func canFinish(numCourses int, prereqs [][]int) bool {
    graph := make([][]int, numCourses)
    for _, prereq := range prereqs {
        to, from := prereq[0], prereq[1]
        graph[from] = append(graph[from], to)
    }
    seen := make([]int, numCourses)
    for i := 0; i < numCourses; i++ {
        if !dfs(i, graph, seen) { return false }
    }
    return true
}

func dfs(course int, graph [][]int, seen []int) bool {
    if seen[course] == 2 {
        return true
    } else if seen[course] == 1 {
        return false
    }
    seen[course] = 1
    for _, next := range graph[course] {
        if !dfs(next, graph, seen) { return false }
    }
    seen[course] = 2
    return true
}

---

func canFinish(numCourses int, prereqs [][]int) bool {
    graph := buildGraph(numCourses, prereqs)
    visited := make([]state, numCourses)
    for course := range graph {
        if !canFinishHelper(course, graph, visited) { return false }
    }
    return true
}

func canFinishHelper(course int, graph [][]int, visited []state) bool {
    if visited[course] == partial {
        return false
    } else if visited[course] == complete {
        return true
    }
    visited[course] = partial
    for _, dependentCourse := range graph[course] {
        if !canFinishHelper(dependentCourse, graph, visited) { return false }
    }
    visited[course] = complete
    return true
}

type state int
const (initial state = iota; partial; complete)

func buildGraph(numCourses int, prereqs [][]int) [][]int {
    graph := make([][]int, numCourses)
    for i := 0; i < len(prereqs); i++ {
        graph[prereqs[i][1]] = append(graph[prereqs[i][1]], prereqs[i][0])
    }
    return graph
}
