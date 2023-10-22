/*
 - "practice", "makes", "perfect", "coding", "makes"
       0          1         2         3         4  
       
   practice = [0]
   makes    = [1]
   perfect  = [2]
   coding   = [3]
   makes    = [4]
   
 - "coding", "practice"
      [3]       [0]
       i         j
       
 if index1[i] < index2[j] {
    res = min(res, index2[j] - index1[i])
    i++
 } else {
    res = min(res, index1[i] - index2[j]) 
 }
 */
type WordDistance struct {
    indexes map[string][]int
    count int
}

func Constructor(wordsDict []string) WordDistance {
    indexes := make(map[string][]int)
    for i, word := range wordsDict {
        indexes[word] = append(indexes[word], i)
    }
    return WordDistance{indexes: indexes, count: len(wordsDict)}
}

func (d *WordDistance) Shortest(word1, word2 string) int {
    index1 := d.indexes[word1]
    index2 := d.indexes[word2]
    var i, j int
    res := d.count
    for i < len(index1) && j < len(index2) {
        if index1[i] < index2[j] {
            res = min(res, index2[j] - index1[i])
            i++
        } else {
            res = min(res, index1[i] - index2[j]) 
            j++
        }
    }
    return res
}

func min(x, y int) int { if x < y { return x }; return y }
