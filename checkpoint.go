package main

import (
	"fmt"
	"sync"
	"time"
)
func worker(id int,checkpoint chan bool,resume chan bool,wg*sync.WaitGroup){
	defer wg.Done()

	fmt.Printf("worker %d:starting\n",id)
	time.Sleep(time.Duration(id)*time.Second)
    
	fmt.Printf("worker %d:checkout reached\n",id)
	checkpoint <- true
	<-resume
	fmt.Printf("worker %d: Resuming\n",id)



}

func main(){
	numworkers:=5
	checkpoint:=make(chan bool)
	resume:=make(chan bool)
	var wg sync.WaitGroup

	for i:=1;i<numworkers;i++{
		wg.Add(1)
		go worker(i,checkpoint,resume,&wg)
	}
   
	for i:=1;i<numworkers;i++{
		<-checkpoint
	}

	fmt.Println("All workers reached checkpoint")
	fmt.Println("Resuming all workers")

	for i:=1;i<numworkers;i++{
		resume<-true

	}

	wg.Wait()
    fmt.Println("All workers completed their work")
}