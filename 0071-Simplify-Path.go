func simplifyPath(path string) string {
    s := list.New()
    tokens := strings.Split(path, "/")
    for _, token := range tokens {
        if token == "" || token == "." {
            continue
        } else if token == ".." {
            if s.Len() != 0 {
                s.Remove(s.Back())
            }
        } else {
            s.PushBack(token)
        }
    }
    if s.Len() == 0 {
        return "/"
    }
    var res []byte
    for s.Len() != 0 {
        res = append(res, '/')
        res = append(res, s.Remove(s.Front()).(string)...)
    }
    return string(res)
}
