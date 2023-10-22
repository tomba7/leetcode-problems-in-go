/*
    "()[]{}"
    "([]){"
         ^
 stack: (, 
 - Iterate the string
 - for every open, push it onto a stack S
 - for every close, peek S and see if it matches the open bracket type
 - If not, return false
 - When done iterating, return false
*/
func isValid(s string) bool {
    stack := list.New()
    for i := 0; i < len(s); i++ {
        if s[i] == '(' || s[i] == '{' || s[i] == '[' {
            stack.PushBack(s[i])
        } else {
            if stack.Len() == 0 { return false }
            opening := stack.Remove(stack.Back()).(byte)
            if s[i] == ')' && opening != '(' ||
                s[i] == '}' && opening != '{'||
                s[i] == ']' && opening != '[' {
                return false
            }
        }
    }
    return stack.Len() == 0
}

---

func isValid(s string) bool {
    stack := []byte{}
    for i := 0; i < len(s); i++ {
        b := s[i]
        if b == '{' || b == '[' || b == '(' {
            stack = append(stack, b)
            continue;
        }
        if len(stack) == 0 { return false }
        top := stack[len(stack) - 1]
        switch b {
        case '}':
            if top != '{' { return false }
        case ']':
            if top != '[' { return false }
        case ')':
            if top != '(' { return false }
        default:
            // If the input character/byte is invalid i.e
            // not equal to "{}[]()", return false.
            return false
        }
        stack = stack[:len(stack) - 1]
    }
    return len(stack) == 0
}
