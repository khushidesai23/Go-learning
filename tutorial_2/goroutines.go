package main
import (
	"fmt"
	"time"
	"sync"
)

//var m = sync.Mutex{}
var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"data1", "data2", "data3", "data4", "data5"}
var results = []string{}

func main(){
	t0 := time.Now()
	for i:=0; i<len(dbData); i++{
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("Total time taken: %v\n", time.Since(t0))
	fmt.Printf("Results: %v\n", results)
}

func dbCall(i int){
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string){
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log(){
	m.RLock()
	fmt.Printf("The current results are: %v\n", results)
	m.RUnlock()
}