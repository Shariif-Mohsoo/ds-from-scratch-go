package operators

import "fmt"

// Days of the week
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// User roles
const (
	Admin       = 10
	Manager     = 20
	Constractor = 30
	Member      = 40
	Guest       = 50
)

func accessGranted() {
	fmt.Println("Granted")
}

func accessDenied() {
	fmt.Println("Denied")
}

func weekDay(day int) bool {
	return day <= 4
}

func Exercise() {
	// The day and role. Change these to check your work.
	// today, role := Tuesday, Guest
	today, role := Sunday, Constractor

	if role == Admin || role == Manager {
		//Access at any time
		accessGranted()
	} else if role == Constractor && !weekDay(today) {
		//Access at weekend
		accessGranted()
	} else if role == Member && weekDay(today) {
		//Access weekDays
		accessGranted()
	} else if role == Guest && (today == Monday || today == Wednesday || today == Friday) {
		accessGranted()
	} else {
		accessDenied()
	}

}
