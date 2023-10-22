func threeSumSmaller(nums []int, target int) int {
    var sum int
    sort.Ints(nums)
    for i := 0; i < len(nums)-2; i++ {
        sum += twoSumSmaller(nums, i+1, target-nums[i])
    }
    return sum
}

func twoSumSmaller(nums []int, i int, target int) int {
    var sum int
    lo, hi := i, len(nums)-1
    for lo < hi {
        if nums[lo] + nums[hi] < target {
            sum += hi - lo
            lo++
        } else {
            hi--
        }
    }
    return sum
}

---

func threeSumSmaller(nums []int, target int) int {
    var result int
    sort.Ints(nums)
    for i := 0; i < len(nums) - 2; i++ {
        j, k := i + 1, len(nums) - 1
        for j < k {
            if nums[i] + nums[j] + nums[k] < target {
                result += k - j
                j++
            } else {
                k--
            }
        }
    }
    return result
}

---

func threeSumSmaller(nums []int, target int) int {
    sort.Ints(nums)
    var total int
    for i := 0; i < len(nums)-2; i++ {
        newTarget := target - nums[i]
        for j, k := i+1, len(nums)-1; j < k; {
            sum := nums[j] + nums[k]
            if sum < newTarget {
                total += k - j
                j++
            } else {
                k--
            }
        }
    }
    return total
}
