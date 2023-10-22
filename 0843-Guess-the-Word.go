func findSecretWord(wordList []string, master *Master) {
    for i := 0; i < 10; i++ {
        index := rand.Intn(len(wordList))
        x := master.Guess(wordList[index])
        if x == 6 { return }
        var newWordList []string
        for _, word := range wordList {
            if word != wordList[index] &&
               matches(word, wordList[index]) == x {
                newWordList = append(newWordList, word)
            }
        }
        wordList = newWordList
    }
}

func matches(word1, word2 string) int {
    var total int
    for i := 0; i < len(word1); i++ {
        if word1[i] == word2[i] { total++ }
    }
    return total
}

---

func findSecretWord(wordList []string, master *Master) {
    rand.Shuffle(len(wordList), func(i, j int) {
        wordList[i], wordList[j] = wordList[j], wordList[i]
    })
    for i := 0; i < 10; i++ {
        guess := wordList[0]
        x := master.Guess(guess)
        if x == 6 { return }
        var newWordList []string
        for _, word := range wordList {
            if match(guess, word) == x {
                newWordList = append(newWordList, word)
            }
            wordList = newWordList
        }
    }
}

func match(word1, word2 string) int {
    var res int
    for i := 0; i < len(word1); i++ {
        if word1[i] == word2[i] {
            res++
        }
    }
    return res
}

---

func findSecretWord(wordList []string, master *Master) {
    matrix := newMatchMatrix(wordList)
    for i := 0; i < 10; i++ {
        index := rand.Intn(len(wordList))
        x := master.Guess(wordList[index])
        if x == 6 { return }
        var newWordList []string
        for j := 0; j < len(wordList); j++ {
            if matrix[index][j] == x {
                newWordList = append(newWordList, wordList[j])
            }
        }
        wordList = newWordList
        matrix = newMatchMatrix(wordList)
    }
}

func newMatchMatrix(wordList []string) [][]int {
    matrix := make([][]int, len(wordList))
    for i := range matrix {
        matrix[i] = make([]int, len(wordList))
    }
    for i := 0; i < len(wordList); i++ {
        for j := 0; j < len(wordList); j++ {
            matrix[i][j] = matches(wordList[i], wordList[j])
        }
    }
    return matrix
}

func matches(word1, word2 string) int {
    var total int
    for i := 0; i < len(word1); i++ {
        if word1[i] == word2[i] { total++ }
    }
    return total
}
