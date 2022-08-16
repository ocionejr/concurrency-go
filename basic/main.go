package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("1")
		time.Sleep(100 * time.Millisecond)
	}()

	go func() {
		defer wg.Done()
		fmt.Println("2")
		time.Sleep(100 * time.Millisecond)
	}()

	go func() {
		defer wg.Done()
		fmt.Println("3")
		time.Sleep(100 * time.Millisecond)
	}()

	wg.Wait()
}
