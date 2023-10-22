// Given an arbitrary ransom note string and another string containing letters
// from all the magazines, write a function that will return true if the ransom
// note can be constructed from the magazines ; otherwise, it will return false.
//
// Each letter in the magazine string can only be used once in your ransom note.
//
// Note:
// You may assume that both strings contain only lowercase letters.
//
// canConstruct("a", "b") -> false
// canConstruct("aa", "ab") -> false
// canConstruct("aa", "aab") -> true

func canConstruct(ransomNote string, magazine string) bool {
    M := make([]int, 27)
    for i := 0; i < len(magazine); i++ { M[magazine[i] - 'a']++ }
    for i := 0; i < len(ransomNote); i++ {
        if M[ransomNote[i] - 'a'] <= 0 { return false }
        M[ransomNote[i] - 'a']--
    }
    return true
}
