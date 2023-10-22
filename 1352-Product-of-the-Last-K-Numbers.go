type ProductOfNumbers struct {
    numbers []int
}

func Constructor() ProductOfNumbers {
    return ProductOfNumbers{
        numbers: []int{1},
    } 
}

func (p *ProductOfNumbers) Add(num int)  {
    n := len(p.numbers)
    if num > 0 {
        p.numbers = append(p.numbers, p.numbers[n-1]*num)
    } else {
        p.numbers = []int{1}
    }
}

func (p *ProductOfNumbers) GetProduct(k int) int {
    n := len(p.numbers)
    if k < n {
        return p.numbers[n-1] / p.numbers[n-k-1]
    }
    return 0
}

---

// Brute force
type ProductOfNumbers struct {
    history []int
}

func Constructor() ProductOfNumbers {
    return ProductOfNumbers{} 
}

func (this *ProductOfNumbers) Add(num int)  {
    this.history = append(this.history, num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
    start := len(this.history) - k
    var res = 1
    for j := start; j < len(this.history); j++ {
        res *= this.history[j]
    }
    return res
}
