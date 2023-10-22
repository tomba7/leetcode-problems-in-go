func twoSum(nums []int, target int) []int {
    indexMap := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        if j, found := indexMap[target - nums[i]]; found {
            return []int{j, i}
        }
        indexMap[nums[i]] = i
    }
    return nil
}

---

func twoSum(nums []int, target int) []int {
    indexMap := make(map[int]int)
    for j, n := range nums {
        if i, ok := indexMap[target-n]; ok {
            return []int{i, j}
        }
        indexMap[n] = j
    }
    return nil
}
