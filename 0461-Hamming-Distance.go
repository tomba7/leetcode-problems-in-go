func hammingDistance(x, y int) int {
    xor := x ^ y
    distance := 0
    for xor != 0 {
        xor &= xor - 1
        distance++
    }
    return distance
}
