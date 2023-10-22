type TwoSum struct {
    freq map[int]int
}


/** Initialize your data structure here. */
func Constructor() TwoSum {
    return TwoSum{freq: map[int]int{}}
}


/** Add the number to an internal data structure.. */
func (t *TwoSum) Add(number int)  {
    t.freq[number]++
}


/** Find if there exists any pair of numbers which sum is equal to the value. */
func (t *TwoSum) Find(value int) bool {
    for num, count := range t.freq {
        x := value - num
        if t.freq[x] != 0 && (num != x || count > 1) {
            return true
        }
    }
    return false
}


/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
