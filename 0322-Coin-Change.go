func coinChange(coins []int, amount int) int {
	table := make([]int, amount+1)
	table[0] = 0
	for a := 1; a <= amount; a++ {
		table[a] = math.MaxInt32
	}
	for a := 1; a <= amount; a++ {
		for _, coin := range coins {
			if coin > a { continue }
			table[a] = min(table[a], table[a-coin])
		}
		if table[a] != math.MaxInt32 { table[a]++ }
	}
	if table[amount] == math.MaxInt32 { return -1 }
	return table[amount]
}

func min(x, y int) int {
    if x < y { return x }; return y
}
