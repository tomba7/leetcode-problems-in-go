type LFUCache struct {
    index      map[int]*keyValue
    listByFreq map[int]*list.List
    minFreq    int
    capacity   int
}

type keyValue struct {
    key, value, freq int
    element *list.Element
}

func Constructor(capacity int) LFUCache {
    cache := LFUCache{
        index: make(map[int]*keyValue),
        listByFreq:  make(map[int]*list.List),
        capacity: capacity,
    }
    cache.listByFreq[1] = list.New()
    return cache
}

func (c *LFUCache) update(kv *keyValue) {
    c.listByFreq[kv.freq].Remove(kv.element)
    if kv.freq == c.minFreq && c.listByFreq[kv.freq].Len() == 0 {
        c.minFreq++
    }
    kv.freq++
    if _, exist := c.listByFreq[kv.freq]; !exist {
        c.listByFreq[kv.freq] = list.New()
    }
    kv.element = c.listByFreq[kv.freq].PushBack(kv.key)
}

func (c *LFUCache) Get(key int) int {
    kv, contains := c.index[key]
    if !contains { return -1 }
    c.update(kv)
    return kv.value
}

func (c *LFUCache) Put(key int, value int)  {
    if c.capacity == 0 { return }
    kv, contains := c.index[key]
    if contains {
        kv.value = value
        c.update(kv)
        return
    }
    if len(c.index) == c.capacity {
        list := c.listByFreq[c.minFreq]
        keyToRemove := list.Remove(list.Front()).(int)
        delete(c.index, keyToRemove)
    }
    c.minFreq = 1
    c.index[key] = &keyValue{
        key: key, value: value, freq: 1, element: c.listByFreq[1].PushBack(key),
    }
}
