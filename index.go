package main

import (
	"fmt"
	"time"
	// "sync"
)

func main(){
	// var wg sync.WaitGroup
	messageChan := make(chan string)
	// wg.Add(1)
	go func ()  {
		// defer wg.Done()
		time.Sleep(2*time.Second)
		messageChan <- "Hello from goroutine"
	}()
	fmt.Println("Hello from main 1")
	message := <-messageChan
	fmt.Println("Hello from main 2")
	fmt.Println(message)

	// wg.Wait()
	close(messageChan)
}
