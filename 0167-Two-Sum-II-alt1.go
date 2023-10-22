func twoSum(numbers []int, target int) []int {
    i, j := 0, len(numbers) - 1
    for i < j {
        sum := numbers[i] + numbers[j]
        if target < sum {
            j--
        } else if target > sum {
            i++
        } else {
            return []int{i + 1, j + 1}
        }
    }
    return nil
}
