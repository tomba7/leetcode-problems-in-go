type MyCircularQueue struct {
    queue []int
    front, tail int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{
        queue: make([]int, k),
        front: -1,
        tail: -1,
    } 
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.IsFull() { return false } 
    if this.front == -1 { this.front = 0 }
    this.tail = (this.tail+1) % len(this.queue)
    this.queue[this.tail] = value 
    return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
    if this.IsEmpty() { return false } 
    this.queue[this.front] = 0
    if this.front == this.tail {
        this.front, this.tail = -1, -1
    } else {
        this.front = (this.front+1) % len(this.queue)
    }
    return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
    if this.IsEmpty() { return -1 } 
    return this.queue[this.front]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
    if this.IsEmpty() { return -1 } 
    return this.queue[this.tail]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
    return this.front == -1
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
    nextTail := (this.tail+1) % len(this.queue)
    return nextTail == this.front
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
