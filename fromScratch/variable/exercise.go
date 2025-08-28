package variable

import "fmt"

func Exercise() {
	var favColor = "black"
	fmt.Println("My fav color:", favColor)

	birthYear, ageInYears := 1987, 34
	fmt.Println("BirthYear:", birthYear, "\nageInYear:", ageInYears, "\n")

	var (
		firstInitial = 'M'
		lastInitial  = 'M'
	)
	fmt.Println("Using ruins")
	fmt.Println("firstInitial:", firstInitial, "\nlastInitial:", lastInitial, "\n")

	var ageInDays int
	ageInDays = 365 * ageInYears
	fmt.Println("Age in days:", ageInDays)
}
