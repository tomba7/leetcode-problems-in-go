func readBinaryWatch(num int) []string {
    result := []string{}
    for hour := 0; hour < 12; hour++ {
        for minute := 0; minute < 60; minute++ {
            if countBits(hour << 6 | minute) == num {
                result = append(result, fmt.Sprintf("%d:%02d", hour, minute))
            }
        }
    }
    return result
}

func countBits(x int) int {
    count := 0
    for x != 0 {
        x &= x - 1
        count++
    }
    return count
}
