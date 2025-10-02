package RWMutex


import (
	"fmt"
	"sync"
	)

var (
	data = 0
	rwMux  sync.RWMutex
	wg   sync.WaitGroup
   )

func readData(id int){
	defer wg.Done()
	rwMux.RLock()
	fmt.Printf("\nReader %d: data = %d\n",id,data)
	rwMux.RUnlock()
}
func writeData(val int){
	defer wg.Done()
	rwMux.Lock()
	data = val
	fmt.Printf("Write: new data = %d\n",data)
	rwMux.Unlock()
}

func RWMutexBasics(){
	wg.Add(5)
	
	go readData(1)
	go readData(2)
	go writeData(105)
	go readData(3)
	go readData(4)
	
	wg.Wait()
}
