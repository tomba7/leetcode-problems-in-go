/*
 - Split("."), len(tokens) == 4, check for IP4
   - no leading zeros
   - 0 <= x <= 255
 - Split(":"), len(tokens) == 6, check for IP6
   - 1 <= len(token) <= 4
   - digits
   - 'a' <= ch <= 'f'
   - 'A' <= ch <= 'F'
 - Neither
 */
func validIPAddress(queryIP string) string {
    if strings.Count(queryIP, ".") == 3 {
        return validIPV4Address(queryIP)
    } else if strings.Count(queryIP, ":") == 7 {
        return validIPV6Address(queryIP)
    }
    return "Neither"
}

func validIPV4Address(queryIP string) string {
    tokens := strings.Split(queryIP, ".")
    if len(tokens) != 4 { return "Neither" }
    for _, token := range tokens {
        if len(token) > 1 && token[0] == '0' { return "Neither" }
        num, err := strconv.Atoi(token)
        if err != nil || num < 0 || num > 255 {
            return "Neither"
        }
    }
    return "IPv4"
}

func validIPV6Address(queryIP string) string {
    tokens := strings.Split(queryIP, ":")
    if len(tokens) != 8 { return "Neither" }
    for _, token := range tokens {
        if len(token) < 1 || len(token) > 4 { return "Neither" }
        for i := 0; i < len(token); i++ {
            ch := token[i]
            if !digit(ch) && !letter(ch) { return "Neither" }
        }
    }
    return "IPv6"
}

func digit(ch byte) bool { return '0' <= ch && ch <= '9' }
func letter(ch byte) bool { return 'a' <= ch && ch <= 'f' || 'A' <= ch && ch <= 'F' }

---

func validIPAddress(queryIP string) string {
    for i := 0; i < len(queryIP); i++ {
        if queryIP[i] == '.' {
            return validIPV4Address(queryIP)
        } else if queryIP[i] == ':' {
            return validIPV6Address(queryIP)
        }
    }
    return "Neither"
}

func validIPV4Address(queryIP string) string {
    tokens := strings.Split(queryIP, ".")
    if len(tokens) != 4 { return "Neither" }
    for _, token := range tokens {
        if len(token) > 1 && token[0] == '0' { return "Neither" }
        num, err := strconv.Atoi(token)
        if err != nil || num < 0 || num > 255 {
            return "Neither"
        }
    }
    return "IPv4"
}

func validIPV6Address(queryIP string) string {
    tokens := strings.Split(queryIP, ":")
    if len(tokens) != 8 { return "Neither" }
    for _, token := range tokens {
        if len(token) < 1 || len(token) > 4 { return "Neither" }
        for i := 0; i < len(token); i++ {
            ch := token[i]
            if !digit(ch) && !letter(ch) { return "Neither" }
        }
    }
    return "IPv6"
}

func digit(ch byte) bool { return '0' <= ch && ch <= '9' }
func letter(ch byte) bool { return 'a' <= ch && ch <= 'f' || 'A' <= ch && ch <= 'F' }
