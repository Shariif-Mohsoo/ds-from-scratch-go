package main

import (
	"fmt"
	"sync"
	"time"
)

// Synchronize go routines.
var wg = sync.WaitGroup{}

func goRoutines() {
	fmt.Println("Welcome to explore go routines")
	//first write the blocking code
	blocked()
	//after block
	fmt.Println("Control comes out of block function")

	//now concurrency
	//create a thread using go keyword (code for this will run in the background)
	wg.Add(1) //add total number of goroutines like 1 ,2 or n; sets counter = 1,2 or n
	go blocking()
	fmt.Println("Concurrency has been achived successfully")

	//it will check whether the all go routines are finished or not; if not then wait.
	wg.Wait() //checks counter = 0
}

func blocked() {
	//using time package
	time.Sleep(10 * time.Second)
	fmt.Print("\nWe are in block function\n")
}

func blocking() {
	//using time package
	time.Sleep(10 * time.Second)
	fmt.Print("\nWe are in block function\n")

	//it will decrement the go routine (size n ) by 1
	defer wg.Done() // always use defer to ensure it runs  //decrement the counter by 1
}
