package interfaces

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("Cook Chicken")
}

func (s Salad) PrepareDish() {
	fmt.Println("Chop Salad")
	fmt.Println("Add dressing")
}

func PrepareDishes(dishes []Preparer) {
	fmt.Println("Prepare dishes:")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("--Dish: %v--\n", dish)
		dish.PrepareDish()
	}
	fmt.Println()
}

func Demo() {
	dishes := []Preparer{Chicken("Chicken Wings"), Salad("Iceberg Salad")}
	PrepareDishes(dishes)
}
