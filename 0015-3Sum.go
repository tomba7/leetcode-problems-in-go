func threeSum(nums []int) [][]int {
    var res [][]int
    sort.Ints(nums)
    for i := 0; i < len(nums)-2; i++ {
        if i != 0 && nums[i] == nums[i-1] { continue }
        target := -nums[i]
        j, k := i+1, len(nums)-1
        for j < k {
            if j != i+1 && nums[j] == nums[j-1] {
                j++
                continue
            }
            if k != len(nums)-1 && nums[k] == nums[k+1] {
                k--
                continue
            }
            sum := nums[j] + nums[k]
            if sum < target {
                j++
            } else if sum > target {
                k--
            } else {
                res = append(res, []int{nums[i], nums[j], nums[k]})
                j++
                k--
            }
        }
    }
    return res
}

---

func threeSum(nums []int) (result [][]int) {
    sort.Ints(nums)
    for i := 0; i < len(nums); i++ {
        if i != 0 && nums[i] == nums[i - 1] { continue }
        target := -nums[i]
        j, k := i + 1, len(nums) - 1
        for j < k {
            sum := nums[j] + nums[k]
            if sum > target {
                k--
            } else if sum < target {
                j++
            } else {
                result = append(result, []int{nums[i], nums[j], nums[k]})
                for ; j < k && nums[j] == nums[j + 1]; j++ {}
                for ; j < k && nums[k] == nums[k - 1]; k-- {}
                j, k = j + 1, k - 1
            }
        }
    }
    return
}

---

// This is a lot slower ~ 1300 ms (compared to 112 ms for the solution above)
func threeSum(arr []int) [][]int {
    var res [][]int
    if len(arr) < 3 { return res }
    sort.Ints(arr)
    m := map[string]struct{}{}
    for i := 0; i < len(arr); i++ {
        t := -arr[i]
        for j, k := i+1, len(arr)-1; j < k; {
            s := arr[j] + arr[k]
            if s > t {
                k--
            } else if s < t {
                j++
            } else {
                key := getKey(arr, i, j, k)
                if _, ok := m[key]; !ok {
                    res = append(res, []int{arr[i], arr[j], arr[k]})
                    m[key] = struct{}{}
                }
                k--
                j++
            }
        }
    }
    return res
}

func getKey(arr []int, i, j, k int) string {
    return fmt.Sprintf("%d,%d,%d", arr[i], arr[j], arr[k])
}
