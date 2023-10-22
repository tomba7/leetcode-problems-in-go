type Logger struct {
    table map[string]int
}

/** Initialize your data structure here. */
func Constructor() Logger {
    return Logger{table: map[string]int{}}
}

/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
        If this method returns false, the message will not be printed.
        The timestamp is in seconds granularity. */
func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool {
    lastTimestamp, ok := l.table[message]
    if !ok || timestamp - lastTimestamp >= 10 {
        l.table[message] = timestamp
        return true
    }
    return false
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
