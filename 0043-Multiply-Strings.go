func multiply(num1, num2 string) string {
	cache := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			product := int(num1[i]-'0')*int(num2[j]-'0') + cache[i+j+1]
			cache[i+j+1] = product % 10
			cache[i+j] += product / 10
		}
	}
	var res []byte
	skip := true
	for i := 0; i < len(cache); i++ {
		if cache[i] == 0 && skip {
			continue
		}
		skip = false
		res = append(res, byte(cache[i])+'0')
	}
	if len(res) == 0 {
		return "0"
	}
	return string(res)
}
