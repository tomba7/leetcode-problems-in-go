const (
    windowSize = 300
)

type HitCounter struct {
    window *list.List
}

func Constructor() HitCounter {
    return HitCounter{
        window: list.New(),
    }    
}

func (c *HitCounter) Hit(timestamp int)  {
    c.window.PushBack(timestamp)
}


func (c *HitCounter) GetHits(timestamp int) int {
    for c.window.Len() != 0 && timestamp - c.window.Front().Value.(int) >= windowSize {
        c.window.Remove(c.window.Front())
    }
    return c.window.Len()
}

---

const (
    windowSize = 300
)

type Hit struct {
    timestamp, count int
}

type HitCounter struct {
    window *list.List
    total  int
}

func Constructor() HitCounter {
    return HitCounter{
        window: list.New(),
    }
}

func (c *HitCounter) Hit(timestamp int)  {
    if c.window.Len() == 0 || c.window.Back().Value.(*Hit).timestamp != timestamp {
        c.window.PushBack(&Hit{
            timestamp: timestamp, count: 1,
        })
    } else {
        c.window.Back().Value.(*Hit).count++
    }
    c.total++
}

func (c *HitCounter) GetHits(timestamp int) int {
    for c.window.Len() != 0 && timestamp - c.window.Front().Value.(*Hit).timestamp >= windowSize {
        c.total -= c.window.Front().Value.(*Hit).count
        c.window.Remove(c.window.Front())
    }
    return c.total
}
