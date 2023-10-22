/*
 - Whenever we need our index (lo) to go out bounds (beyond n-1)
   to get the desired result, we choose hi as 'n' instead of 'n-1'
   because we can end up 'lo' pointing to 'hi' and that is what we want
 - In this problem, if the letters are ['c', 'f', j'] and target = 'j',
   then when letters[mid] == target ('j'), then we want our 'lo'
   to point to 3, so that we can do j % n = 3 % 3 = 0. Because we
   want our result to be 'a'.
 */
func nextGreatestLetter(letters []byte, target byte) byte {
    n := len(letters)
    lo, hi := 0, n
    for lo < hi {
        mid := lo + (hi-lo)/2
        if target >= letters[mid] {
            lo = mid + 1
        } else {
            hi = mid
        }
    }
    return letters[lo % n]
}
