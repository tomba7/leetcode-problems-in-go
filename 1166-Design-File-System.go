type Node struct {
    Value    int
    Children map[string]*Node
}

func NewNode() *Node {
    return &Node{
        Value:    -1,
        Children: make(map[string]*Node),
    }
}

type FileSystem struct {
    root *Node
}

func Constructor() FileSystem {
    return FileSystem{root: NewNode()} 
}

func (fs *FileSystem) CreatePath(path string, value int) bool {
    if path == "" || path == "/" { return false }
    tokens := strings.Split(path, "/")
    curr := fs.root
    for i := 1; i < len(tokens); i++ {
        child, exist := curr.Children[tokens[i]]
        if !exist {
            if i == len(tokens)-1 {
                child = NewNode()
                curr.Children[tokens[i]] = child
            } else {
                return false
            }
        }
        curr = child
    }
    if curr.Value != -1 { return false }
    curr.Value = value
    return true 
}

func (fs *FileSystem) Get(path string) int {
    if path == "" || path == "/" { return -1 }
    tokens := strings.Split(path, "/")
    curr := fs.root
    for i := 1; i < len(tokens); i++ {
        child, exist := curr.Children[tokens[i]]
        if !exist { return -1 }
        curr = child
    }
    return curr.Value 
}

---

type Node struct {
    Value    int
    Children map[string]*Node
}

func NewNode() *Node {
    return &Node{
        Value:    -1,
        Children: make(map[string]*Node),
    }
}

type FileSystem struct {
    root *Node
}

func Constructor() FileSystem {
    return FileSystem{root: NewNode()} 
}

func (fs *FileSystem) CreatePath(path string, value int) bool {
    if path == "" || path == "/" { return false }
    tokens := strings.Split(path, "/")
    curr := fs.root
    for i := 1; i < len(tokens); i++ {
        child, exist := curr.Children[tokens[i]]
        if !exist {
            if i == len(tokens)-1 {
                child = NewNode()
                curr.Children[tokens[i]] = child
            } else {
                return false
            }
        } else {
            if i == len(tokens)-1 { return false }
        }
        curr = child
    }
    curr.Value = value
    return true 
}

func (fs *FileSystem) Get(path string) int {
    if path == "" || path == "/" { return -1 }
    tokens := strings.Split(path, "/")
    curr := fs.root
    for i := 1; i < len(tokens); i++ {
        child, exist := curr.Children[tokens[i]]
        if !exist { return -1 }
        curr = child
    }
    return curr.Value 
}
