package arrays

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanLiness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "is clean")
		} else {
			fmt.Println(room.name, "is dirty")
		}
	}
}

func Demo() {
	rooms := [...]Room{
		{name: "Office", cleaned: true},
		{name: "Warehouse"},
		{name: "Reception"},
		{name: "Ops", cleaned: true},
	}
	checkCleanLiness(rooms)
}
