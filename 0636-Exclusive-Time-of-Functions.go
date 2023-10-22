/*
    n = 2,
    logs = ["0:start:0","1:start:2","1:end:5","0:end:6"]
    Output: [3,4]
    
      0 1 2 3 4 5 6
    1     _ _ _ _
    0 _ _         _
    
    (1, 2)
    (0, 0)
    0: 2
    1: 4
    
    ["0:start:0","0:start:2","0:end:5","0:start:6","0:end:6","0:end:7"]
    
      0 1 2 3 4 5 6 7
    0             _
    0     _ _ _ _
    0 _ _           _
    
    
    ["0:start:0","0:start:2","0:end:5","1:start:6","1:end:6","0:end:7"]
    
      0 1 2 3 4 5 6 7
    1             _
    0     _ _ _ _
    0 _ _           _
    
    ["0:start:0","0:start:2","0:end:5","1:start:7","1:end:7","0:end:8"]
    
      0 1 2 3 4 5 6 7 8
    1               _
    0     _ _ _ _
    0 _ _         _   _
*/
func exclusiveTime(n int, logs []string) []int {
    res := make([]int, n)
    q := list.New()
    for i := 0; i < len(logs); i++ {
        curr := newLogEntry(logs[i])
        if curr.start {
            q.PushBack(curr)
        } else {
            start, end := q.Remove(q.Back()).(*logEntry), curr
            res[end.id] += end.timestamp - start.timestamp + 1
            if q.Len() != 0 {
                caller := q.Back().Value.(*logEntry)
                res[caller.id] -= end.timestamp - start.timestamp + 1
            }
        }
    }
    return res
}

type logEntry struct {
    id        int
    timestamp int
    start     bool
}

func newLogEntry(line string) *logEntry {
    tokens := strings.Split(line, ":")
    id, _ := strconv.Atoi(tokens[0])
    timestamp, _ := strconv.Atoi(tokens[2])
    return &logEntry{
        id: id, timestamp: timestamp, start: tokens[1] == "start",
    }
}

---

func exclusiveTime(n int, logs []string) []int {
    res := make([]int, n)
    q := list.New()
    for i := 0; i < len(logs); i++ {
        funcId, timestamp, isStart := logEntry(logs[i])
        if isStart {
            q.PushBack([]int{funcId, timestamp})
        } else {
            start := q.Remove(q.Back()).([]int)
            res[funcId] += timestamp - start[1] + 1
            if q.Len() != 0 {
                top := q.Back().Value.([]int)
                res[top[0]] -= timestamp - start[1] + 1
            }
        }
    }
    return res
}

func logEntry(logEntry string) (int, int, bool) {
    entry := strings.Split(logEntry, ":")
    funcId, _ := strconv.Atoi(entry[0])
    timestamp, _ := strconv.Atoi(entry[2])
    return funcId, timestamp, entry[1] == "start"
}
