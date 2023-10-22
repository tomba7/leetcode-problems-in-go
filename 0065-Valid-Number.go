/*
4:40 PM

[integer]
[decimal]e[integer]
[decimal]E

<<decimal>>
(+)ddd.
(-)ddd.ddd
(+).ddd

<<integer>>
(+)ddd
 */
func isNumber(s string) bool {
    i, n := 0, len(s)
    validDigits := false
    if i < n && (s[i] == '-' || s[i] == '+') { i++ }
    for i < n && digit(s[i]) {
        i++
        validDigits = true
    }
    if i < n && s[i] == '.' { i++ }
    for i < n && digit(s[i]) {
        i++
        validDigits = true
    }
    if i < n && (s[i] == 'e' || s[i] == 'E') && validDigits {
        i++
        validDigits = false
        if i < n && (s[i] == '-' || s[i] == '+') { i++ }
        for i < n && digit(s[i]) {
            i++
            validDigits = true
        }
    }
    return validDigits && i == n
}

func digit(x byte) bool { return '0' <= x && x <= '9' }

---

/*
 - +/- signs are optional
 - When we see digits before or after a '.' we want to mark it seen
 - If a decimal point is seen, we want to make sure that
   there are valid digits either
   (a) before it
   (b) after it
   (c) or both
 - When an 'e' or 'E' is encountered, we want to make sure that
   we have already seen valid digits before we got there
 - After an 'e' or 'E' is seen, we need to verify that valid
   digits are detected
 - Finally at the end, we want to make sure that valid digits
   have been seen/marked AND that we've consumed all of s
*/
func isNumber(s string) bool {
    n := len(s)
    var validDigitsFound bool
    var i int
    // Optional sign
    if i < n && (s[i] == '+' || s[i] == '-') {
        i++
    }
    // Mark that we've seen digits before the .
    for ; i < n && digit(s[i]); i++ {
        validDigitsFound = true
    }
    // If we've seen a decimal point, we expect digits
    // either before or after.
    if i < n && s[i] == '.' { i++ }
    for ; i < n && digit(s[i]); i++ {
        validDigitsFound = true
    }
    // If we get to 'e', we want to verify that valid digits
    // have been seen before
    if i < n && (s[i] == 'e' || s[i] == 'E') && validDigitsFound {
        i++
        validDigitsFound = false
        if i < n && (s[i] == '+' || s[i] == '-') {
            i++
        }
        for ; i < n && digit(s[i]); i++ {
            validDigitsFound = true
        }
    }
    // Verify that we have seen valid digits so far and have
    // succesfully parsed all the characters in s.
    return validDigitsFound && i == n
}

func digit(x byte) bool { return '0' <= x && x <= '9' }
