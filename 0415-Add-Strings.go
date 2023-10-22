func addStrings(num1, num2 string) string {
    var carry byte
    result := []byte{}
    for i, j := len(num1) - 1, len(num2) - 1; i >= 0 || j >= 0 || carry != 0; i, j = i - 1, j - 1 {
        if i >= 0 { carry += num1[i] - '0'}
        if j >= 0 { carry += num2[j] - '0'}
        result = append(result, carry % 10 + '0')
        carry /= 10
    }
    for i, j := 0, len(result) - 1; i < j; i, j = i + 1, j - 1 {
        result[i], result[j] = result[j], result[i]
    }
    return string(result)
}

---

/*  
- Make sure num1 > num2, if not swap
- Set carry to 0
- Starting from the end, extract both digits and sum them + carry
- Convert the sum and append to result, and assign new carry
- Do this till the carry is zero
- Finally reverse the list and convert to string
    456
     77
    ---
     33
 */
func addStrings(num1 string, num2 string) string {
    if len(num1) < len(num2) {
        return addStrings(num2, num1)
    }
    m, n := len(num1), len(num2)
    var carry byte
    var result []byte
    for i, j := m-1, n-1; i >= 0 || j >= 0 || carry != 0; {
        sum := carry
        if i >= 0 {
            sum += num1[i] - '0'
            i--
        }
        if j >= 0 {
            sum += num2[j] - '0'
            j--
        }
        result = append(result, '0' + sum % 10)
        carry = sum / 10
    }
    for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }
    return string(result)
}
