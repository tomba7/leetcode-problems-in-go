/*
10001010 <--- num
10001001
------------
10001000
10000111
------------
10000000
01111111
---------
000000000

num = num & (num-1)

 */
func hammingWeight(num uint32) int {
    var total int
    for num != 0 {
        num &= num - 1
        total++
    }
    return total
}
