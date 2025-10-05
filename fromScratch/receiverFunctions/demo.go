package receiverfunctions

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

// Regular function
func occupySpace(lot *ParkingLot, spaceNum int) {
	//spaceNum will be idx
	lot.spaces[spaceNum-1].occupied = true
}

// Reciever function(pointer)
func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

//Free occupied space

func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func Demo() {
	lot := ParkingLot{spaces: make([]Space, 10)}
	fmt.Println("Initial:", lot)
	lot.occupySpace(1)
	occupySpace(&lot, 2)
	fmt.Println("After occupied:", lot)
	lot.vacateSpace(1)
	fmt.Println("After vacate:", lot)
}
