func isHappy(n int) bool {
    slow, fast := n, n
    for {
        slow = digitSquareSum(slow)
        fast = digitSquareSum(digitSquareSum(fast))
        if slow == fast {
            if slow == 1 { return true }
            return false
        }
    }
}

func digitSquareSum(n int) int {
    var sum int
    for n != 0 {
        sum += (n % 10) * (n % 10)
        n /= 10
    }
    return sum
}

---

func isHappy(n int) bool {
    seen := map[int]bool{}
    for n != 1 {
        if seen[n] { return false }
        seen[n] = true
        n = addDigits(n)
    }
    return true
}

func addDigits(n int) int {
    var sum int
    for n != 0 {
        digit := n % 10
        sum += digit*digit
        n /= 10
    }
    return sum
}

---

func isHappy(n int) bool {
    var slow, fast = n, n
    for {
        slow = squaredSum(slow)
        fast = squaredSum(fast)
        fast = squaredSum(fast)
        if slow == fast { break }
    }
    if slow == 1 { return true }
    return false
}

func squaredSum(n int) int {
    var sum int
    for n != 0 {
        sum += (n%10) * (n%10)
        n /= 10
    }
    return sum
}
