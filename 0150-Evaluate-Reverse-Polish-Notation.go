func evalRPN(tokens []string) int {
    var n int
    stack := []int{}
    for _, token := range tokens {
        if token == "/" || token == "*" || token == "+" || token == "-" {
            x, y := stack[len(stack) - 2], stack[len(stack) - 1]
            stack = stack[:len(stack) - 2]
            switch token {
            case "/":
                n = x / y
            case "*":
                n = x * y
            case "+":
                n = x + y
            case "-":
                n = x - y
            }
        } else {
            n, _ = strconv.Atoi(token)
        }
        stack = append(stack, n)
    }
    return n
}
