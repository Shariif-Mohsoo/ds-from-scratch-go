package main

import "fmt"

// First interface
type Reader interface {
	Read()
}

// Second interface
type Writer interface {
	Write()
}

// Third interface (combined)
type Closer interface {
	Close()
}

// Struct that implements all 3
type File struct {
	FileName string
}

// Implement Reader
func (f File) Read() {
	fmt.Println("Reading from", f.FileName)
}

// Implement Writer
func (f File) Write() {
	fmt.Println("Writing to", f.FileName)
}

// Implement Closer
func (f File) Close() {
	fmt.Println("Closing", f.FileName)
}

// Main function
func multipleInterface() {
	file := File{FileName: "data.txt"}

	var r Reader = file
	var w Writer = file
	var c Closer = file

	r.Read()
	w.Write()
	c.Close()
}
