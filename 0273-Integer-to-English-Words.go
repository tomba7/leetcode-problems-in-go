/*
single digits       = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
less than twenty    = [11, 12, 13, 14, 15, 16, 17, 18, 19]
tens                = [10, 20, 30, 40, 50, 60, 70, 80, 90]

thousands           = "", "Thousand", "Million", "Billion"

3   2   1   0
1,073,741,824

1,234,567
*/
func numberToWords(num int) string {
    if num == 0 { return "Zero" }
    var res string
    var i int
    for num > 0 {
        if num % 1000 != 0 {
            res = numberToWordsHelper(num % 1000) + thousands[i] + " " + res
        }
        num /= 1000
        i++
    }
    return strings.TrimSpace(res)
}

func numberToWordsHelper(num int) string {
    if num == 0 {
        return ""
    } else if num < 20 {
        return lessThanTwenty[num] + " "
    } else if num < 100 {
        return tens[num/10] + " " + numberToWordsHelper(num%10)
    }
    return lessThanTwenty[num/100] + " Hundred " + numberToWordsHelper(num%100)
}

var lessThanTwenty = []string{
    "", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
    "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen",
    "Eighteen", "Nineteen",
}

var tens = []string{
    "", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety",
}

var thousands = []string{"", "Thousand", "Million", "Billion"}
