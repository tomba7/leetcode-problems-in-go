type Event struct {
	start, end  int
	left, right *Event
}

type MyCalendar struct {
	root *Event
	ok   bool
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (c *MyCalendar) search(curr *Event, start int, end int) *Event {
	c.ok = true
	if curr == nil {
		return &Event{start: start, end: end}
	}
	if curr.start >= end {
		curr.left = c.search(curr.left, start, end)
	} else if curr.end <= start {
		curr.right = c.search(curr.right, start, end)
	} else {
		c.ok = false
	}
	return curr
}

func (c *MyCalendar) Book(start int, end int) bool {
	if c.root == nil {
		c.root = &Event{start: start, end: end}
		return true
	}
	c.search(c.root, start, end)
	return c.ok
}

---

type Event struct {
    Start, End int
    Left, Right *Event
}

type MyCalendar struct {
    root *Event
}

func Constructor() MyCalendar {
    return MyCalendar{}
}

func (c *MyCalendar) Book(start, end int) bool {
    if c.root == nil {
        c.root = &Event{Start: start, End: end}
        return true
    }
    return c.book(c.root, start, end)
}

func (c *MyCalendar) book(curr *Event, start, end int) bool {
    if start < curr.End && end > curr.Start {
        return false
    }
    if start >= curr.End {
        if curr.Right == nil {
            curr.Right = &Event{Start: start, End: end}
            return true
        } else {
            return c.book(curr.Right, start, end)
        }
    } else {
        if curr.Left == nil {
            curr.Left = &Event{Start: start, End: end}
            return true
        } else {
            return c.book(curr.Left, start, end)
        }
    }
}
