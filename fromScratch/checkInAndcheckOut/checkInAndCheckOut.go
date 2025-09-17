package checkinandcheckout

import (
	"fmt"
	"time"
)

type Title string //book title
type Name string  // library member name

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int //total owned by the library
	lended int //total books lended out
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkOut.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "though", returnTime)
	}
}

func printMemberAudits(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

func printLibraryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/ lended:", book.lended)
	}
	fmt.Println()
}

func checkoutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}
	if book.lended == book.total {
		fmt.Println("No more books available to lend")
		return false
	}

	book.lended += 1
	library.books[title] = book
	member.books[title] = LendAudit{checkOut: time.Now()}
	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}
	audit, found := member.books[title]
	if !found {
		fmt.Println("Member did not check out this book")
		return false
	}
	book.lended -= 1
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit
	return true
}

func CheckInAndCheckOut() {
	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}
	library.books["Webapps in Go"] = BookEntry{
		total:  4,
		lended: 0,
	}
	library.books["The Little Go Book"] = BookEntry{
		total:  3,
		lended: 0,
	}
	library.books["Let's learn Go"] = BookEntry{
		total:  2,
		lended: 0,
	}
	library.books["Go Bootcamp"] = BookEntry{
		total:  1,
		lended: 0,
	}
	library.members["Jayson"] = Member{"Jayson", make(map[Title]LendAudit)}
	library.members["Billy"] = Member{"Billy", make(map[Title]LendAudit)}
	library.members["Susonna"] = Member{"Susonna", make(map[Title]LendAudit)}
	fmt.Println("\nInitial:")
	printLibraryBooks(&library)
	printMemberAudits(&library)

	member := library.members["Jayson"]
	checkedOut := checkoutBook(&library, "Go Bootcamp", &member)
	fmt.Println("\nCheck out a book:")
	if checkedOut {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}
	returned := returnBook(&library, "Go Bootcamp", &member)
	fmt.Println("\nCheck in a book:")
	if returned {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}
}
