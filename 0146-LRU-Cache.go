type LRUCache struct {
    index    map[int]*list.Element
    elements *list.List
    capacity int
}

type kvPair struct { key, value int }

func Constructor(capacity int) LRUCache {
    return LRUCache{
        index: make(map[int]*list.Element),
        elements: list.New(),
        capacity: capacity,
    }
}

func (c *LRUCache) Get(key int) int {
    elem, exist := c.index[key]
    if !exist {
        return -1
    }
    c.elements.MoveToBack(elem)
    return elem.Value.(*kvPair).value
}

func (c *LRUCache) Put(key int, value int)  {
    elem, exist := c.index[key]
    if !exist {
        elem := c.elements.PushBack(&kvPair{key: key, value: value})
        c.index[key] = elem
    } else {
        elem.Value.(*kvPair).value = value
        c.elements.MoveToBack(elem)
    }
    if c.elements.Len() > c.capacity {
        oldestElem := c.elements.Remove(c.elements.Front())
        delete(c.index, oldestElem.(*kvPair).key)
    }
}

---

type KVPair struct { key, value int }

type LRUCache struct {
    table map[int]*list.Element
    list *list.List
    capacity int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
        table:    make(map[int]*list.Element),
        list:     list.New(),
        capacity: capacity,
    }
}

func (c *LRUCache) Get(key int) int {
    elem, ok := c.table[key]
    if !ok { return -1 }
    c.list.MoveToBack(elem)
    return elem.Value.(KVPair).value
}

func (c *LRUCache) Put(key int, value int)  {
    elem, ok := c.table[key]
    if len(c.table) < c.capacity {
        if ok { c.list.Remove(elem) }
    } else {
        if !ok {
            delete(c.table, c.list.Front().Value.(KVPair).key)
            c.list.Remove(c.list.Front())
        } else {
            c.list.Remove(elem)
        }
    }
    c.table[key] = c.list.PushBack(KVPair{key, value})
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
