func subdomainVisits(cpdomains []string) []string {
    sdTable := map[string]int{}
    for _, cpDomain := range cpdomains {
        i := strings.Index(cpDomain, " ")
        count, _ := strconv.Atoi(cpDomain[:i])
        subDomain := cpDomain[i+1:]
        sdTable[subDomain] += count
        for j := 0; j < len(subDomain); j++ {
            if subDomain[j] == '.' {
                sdTable[subDomain[j+1:]] += count
            }
        }
    }
    var res []string
    for subDomain, count := range sdTable {
        res = append(res, fmt.Sprintf("%d %s", count, subDomain))
    }
    return res
}
