/*
 - For every user, create an array of all the sites he/she visted (along with the time stamp)
   and put it in a hash map
   joe = [(home, 1), (about, 2), (career, 3)]
   james = [(home, 3) and so on....]
 - For every user, sort the array based on the time stamp
 - For every user, create all possible patterns and store the count in a hash map 
 - Maintain maximum "res" as we count
 */
func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
    history := map[string][]*pair{}
    for i := 0; i < len(username); i++ {
        history[username[i]] = append(history[username[i]], &pair{
            website: website[i], timestamp: timestamp[i],
        })
    }
    var res string
    count := map[string]int{}
    for _, pairs := range history {
        sort.Slice(pairs, func(i, j int) bool {
            return pairs[i].timestamp < pairs[j].timestamp
        })
        n := len(pairs)
        seen := map[string]bool{}
        for i := 0; i < n-2; i++ {
            for j := i+1; j < n-1; j++ {
                for k := j+1; k < n; k++ {
                    key := pairs[i].website + " " + pairs[j].website + " " + pairs[k].website
                    if seen[key] { continue }
                    seen[key] = true
                    count[key]++
                    if res == "" || count[key] > count[res] || count[key] == count[res] && key < res {
                        res = key
                    }
                }
            }
        }
    }
    return strings.Split(res, " ")
}

type pair struct { website string; timestamp int }

---

/** 11:20
 [username, website, timestamp]
 
 ["joe", "joe",  "joe",   "james","james","james","james","mary","mary", "mary"]
 ["home","about","career","home", "cart", "maps", "home", "home","about","career"]
 [ 1,     2,      3,       4,      5,      6,      7,      8,     9,      10]

["ua","ua","ua", "ub","ub","ub"]
["a", "b", "a",  "a", "b", "c"]

- Create an array for each user [[website1, ts1], [website2, ts2]] sorted by timestamp
- Then generate all possible patterns (lenght of 3)
  and increment the count of that pattern
    e.g. "home about career" = 2
         "home cart maps"    = 1
         "home cart home"    = 1
         "home maps home"    = 1
         "cart maps home"    = 1
         
         ["home about career"]
         ["home cart maps", "home cart home", "home maps home", "cart maps home"]
        
- Walk through the hash map and create an array sorted by the score
- Get the largest entry, sort and return

 */
func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
    n := len(username)
    visitMap := make(map[string][]Visit)
    for i := 0; i < n; i++ {
        visitMap[username[i]] = append(visitMap[username[i]], Visit{
            website: website[i], timestamp: timestamp[i],
        })
    }
    patternFreq := make(map[string]int)
    for _, visits := range visitMap {
        sort.Slice(visits, func(i, j int) bool {
            return visits[i].timestamp < visits[j].timestamp
        })
        seen := map[string]bool{}
        for i := 0; i < len(visits); i++ {
            for j := i+1; j < len(visits); j++ {
                for k := j+1; k < len(visits); k++ {
                    pattern := fmt.Sprintf("%s %s %s", visits[i].website, visits[j].website, visits[k].website)
                    if seen[pattern] { continue }
                    seen[pattern] = true
                    patternFreq[pattern]++
                }
            }
        }
    }
    var maxScore int
    for _, score := range patternFreq {
        if score > maxScore { maxScore = score }
    }
    var patterns []string
    for pattern, score := range patternFreq {
        if score == maxScore {
            patterns = append(patterns, pattern)
        }
    }
    sort.Strings(patterns)
    res := strings.Split(patterns[0], " ")
    return res
}

type Visit struct {
    website string
    timestamp int
}
