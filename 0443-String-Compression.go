/*
 i
"aabbccc"
 j

- Use two pointers: i to write and j to read.
- j will scan the string from left to right
- at every step we will compare chars[j] against chars[j+1]
- if it is the same we increment the curr char count
- if not then we write the current char at i, increment i
  then write the number at i and increment i

123/10 = 12, 123%10 = 3
12 /10 = 1   12 %10 = 2
1  /10 = 0   1  %10 = 1
  
 */
func compress(chars []byte) int {
    var i, count int
    for j := 0; j < len(chars); j++ {
        count++
        if j+1 == len(chars) || chars[j] != chars[j+1] {
            chars[i] = chars[j]
            i++
            if count > 1 {
                encode(chars, count, &i)
            }
            count = 0
        }
    }
    return i
}

func encode(chars []byte, count int, i *int) {
    var digits []byte
    for count != 0 {
        digits = append(digits, byte(count % 10) + '0')
        count /= 10
    }
    for j := len(digits)-1; j >= 0; j-- {
        chars[*i] = digits[j]
        *i++
    }
}

---

func compress(chars []byte) int {
    var i, j, count int
    for ; j < len(chars); j++ {
        count++
        if j+1 == len(chars) || chars[j] != chars[j+1] {
            chars[i] = chars[j]
            i++
            if count > 1 {
                s := strconv.Itoa(count)
                for j := 0; j < len(s); j++ {
                    chars[i] = s[j]
                    i++
                }
            }
            count = 0
        }
    }
    return i
}

---

func compress(chars []byte) int {
    var i, j, count int
    for ; j < len(chars); j++ {
        count++
        if j+1 == len(chars) || chars[j] != chars[j+1] {
            chars[i] = chars[j]
            i++
            if count > 1 {
                for _, c := range strconv.Itoa(count) {
                    chars[i] = byte(c)
                    i++
                }
            }
            count = 0
        }
    }
    return i
}

---

func compress(chars []byte) int {
    var i, j int
    var count byte
    for ; j < len(chars); j++ {
        count++
        if j+1 == len(chars) || chars[j] != chars[j+1] {
            chars[i] = chars[j]
            i++
            if count > 1 { itoa(chars, &i, count) }
            count = 0
        }
    }
    return i
}

func itoa(chars []byte, i *int, count byte) {
    var nums []byte
    for count != 0 {
        nums = append(nums, count%10 + '0')
        count /= 10
    }
    for j := len(nums)-1; j >= 0; j-- {
        chars[*i] = nums[j]
        *i++
    }
}
