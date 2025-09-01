package diceroller

import (
	"fmt"
	"math/rand"
	"time"
)

func roll(sides int) int {
	// 5 :(0,1,2,3,4) + 1
	// 5: 1,2,3,4,5
	return rand.Intn(sides) + 1
}

func DiceRoller() {
	//initializing random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	//no of dices and sides
	dice, sides := 2, 12
	rolls := 3

	for r := 1; r <= rolls; r++ {
		sum := 0
		for d := 1; d <= dice; d++ {
			rolled := roll(sides)
			sum += rolled
			fmt.Println("Roll #", r, "Die #", d, ":", rolled)
		}
		fmt.Println("Total rolled:", sum)
		switch sum := sum; {
		case sum == 2 && dice == 2:
			fmt.Println("Snake Eyes!")
		case sum == 7:
			fmt.Println("Lucky seven!")
		case sum%2 == 0:
			fmt.Println("Even")
		case sum%2 == 1:
			fmt.Println("Odd")
		}
	}

}
