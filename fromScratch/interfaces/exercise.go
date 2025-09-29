package interfaces

import "fmt"

const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Lift int

type LiftPicker interface {
	PickLift() Lift
}

type Motorcycle string
type Car string
type Truck string

func (m Motorcycle) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(m))
}

func (c Car) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(c))
}

func (t Truck) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(t))
}

// implementing the interface
func (m Motorcycle) PickLift() Lift {
	return SmallLift
}

func (c Car) PickLift() Lift {
	return StandardLift
}

func (t Truck) PickLift() Lift {
	return LargeLift
}

func sendToLift(p LiftPicker) {
	switch p.PickLift() {
	case SmallLift:
		fmt.Printf("send %v to small lift\n", p)
	case StandardLift:
		fmt.Printf("send %v to standard lift\n", p)
	case LargeLift:
		fmt.Printf("send %v to large lift\n", p)
	}
}

func Exercise() {
	car := Car("Sporty")
	truck := Truck("MountainCrusher")
	motorcycle := Motorcycle("Croozer")

	sendToLift(car)
	sendToLift(truck)
	sendToLift(motorcycle)
}
