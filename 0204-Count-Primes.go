func countPrimes(n int) int {
    if n < 2 {return 0}
    table := make([]bool, n + 1)
    table[0], table[1] = true, true
    for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
        if table[i] {continue}
        for j := i * i; j < n; j += i {
            table[j] = true
        }
    }
    numPrimes := 0
    for i := 2; i < n; i++ {
        if !table[i] {
            numPrimes++
        }
    }
    return numPrimes
}

---

func countPrimes(n int) int {
    if n < 2 {return 0}
    table := make([]bool, n + 1)
    table[0], table[1] = true, true
    numPrimes := 0
    for i := 2; i < n; i++ {
        if table[i] {continue}
        numPrimes++
        for j := i * i; j < n; j += i {
            table[j] = true
        }
    }
    return numPrimes
}

---

func countPrimes(n int) int {
    var total int
    table := make([]bool, n)
    for i := 2; i < n; i++ {
        table[i] = true
    }
    for i := 2; i < n; i++ {
        if !table[i] { continue }
        for j := i*2; j < n; j += i {
            table[j] = false
        }
        total++
    }
    return total
}

---

// Brute Force (Time Limit Exceeded)
func countPrimes(n int) int {
    var total int
    for i := 2; i < n; i++ {
        if isPrime(i) { total++ }
    }
    return total
}

func isPrime(n int) bool {
    for x := 2; x < n; x++ {
        if n % x == 0 { return false }
    }
    return true
}

---

// Optimized Brute Force (Time Limit Exceeded)
func countPrimes(n int) int {
    var total int
    for i := 2; i < n; i++ {
        if isPrime(i) { total++ }
    }
    return total
}

func isPrime(n int) bool {
    for x := 2; x <= int(math.Sqrt(float64(n))); x++ {
        if n % x == 0 { return false }
    }
    return true
}

---

// Optimized Brute Force (Time Limit Exceeded)
func countPrimes(n int) int {
    var total int
    for i := 2; i < n; i++ {
        if isPrime(i) { total++ }
    }
    return total
}

func isPrime(n int) bool {
    if n < 2 { return false }
    for x := 2; x*x <= n; x++ {
        if n % x == 0 { return false }
    }
    return true
}
