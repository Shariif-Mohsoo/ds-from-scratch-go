package maps

import "fmt"

func creatingMap() {
	myMap := make(map[string]int)
	myMap = map[string]int{
		"item 1": 1,
		"item 2": 2,
		"item 3": 3,
	}
	fmt.Println(myMap)
	myMap["item 4"] = 4
	fmt.Println(myMap)

	fav := myMap["item 4"]
	fmt.Println(fav)

	delete(myMap, "item 4")
	fmt.Println(myMap)

	price, found := myMap["price"]
	if !found {
		fmt.Println("Price not found")
	} else {
		fmt.Println(price)
	}
	//map iteration
	for key, val := range myMap {
		fmt.Println(key, "=", val)
	}
}

func Maps() {
	creatingMap()
}
