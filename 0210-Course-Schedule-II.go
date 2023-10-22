func findOrder(numCourses int, prereqs [][]int) []int {
    graph := newGraph(numCourses, prereqs)
    seen := make([]int, numCourses)
    var res []int
    for course := 0; course < numCourses; course++ {
        if dfs(graph, seen, course, &res) {
            return nil
        }
    }
    for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
        res[i], res[j] = res[j], res[i]
    }
    return res
}

func dfs(graph [][]int, seen []int, course int, res *[]int) bool {
    if seen[course] == 2 {
        return false
    } else if seen[course] == 1 {
        return true
    }
    seen[course] = 1
    for _, next := range graph[course] {
        if dfs(graph, seen, next, res) { return true }
    }
    *res = append(*res, course)
    seen[course] = 2
    return false
}

func newGraph(numCourses int, prereqs [][]int) [][]int {
    graph := make([][]int, numCourses)
    for _, prereq := range prereqs {
        first, second := prereq[1], prereq[0]
        graph[first] = append(graph[first], second)
    }
    return graph
}

---

func findOrder(numCourses int, prereqs [][]int) []int {
    graph := buildGraph(numCourses, prereqs)
    visited := make([]state, numCourses)
    result, index := make([]int, numCourses), numCourses - 1
    for course := range graph {
        if !findOrderHelper(course, graph, visited, result, &index) { return nil }
    }
    return result
}

func findOrderHelper(course int, graph [][]int, visited []state, result []int, index *int) bool {
    if visited[course] == partial {
        return false
    } else if visited[course] == complete {
        return true
    }
    visited[course] = partial
    for _, dependentCourse := range graph[course] {
        if !findOrderHelper(dependentCourse, graph, visited, result, index) { return false }
    }
    result[*index] = course
    *index--
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
