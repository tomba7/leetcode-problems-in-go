type ParkingSystem struct {
    spaces []int
}

func Constructor(big int, medium int, small int) ParkingSystem {
    spaces := make([]int, 4)
    spaces[1] = big
    spaces[2] = medium
    spaces[3] = small
    return ParkingSystem{spaces: spaces}
}

func (p *ParkingSystem) AddCar(carType int) bool {
    if p.spaces[carType] == 0 {
        return false
    }
    p.spaces[carType]--
    return true
}
