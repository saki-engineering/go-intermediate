package main

import (
	"fmt"
	"sync"
)

var i int
var mu sync.Mutex
var wg sync.WaitGroup

func MyFunc1() {
	defer wg.Done()

	fmt.Println("MyFunc1 start")

	mu.Lock()
	i += 1
	mu.Unlock()

	fmt.Println("MyFunc1 finish")
}

func MyFunc2() {
	defer wg.Done()

	fmt.Println("MyFunc2 start")

	mu.Lock()
	i -= 1
	mu.Unlock()

	fmt.Println("MyFunc2 finish")
}

func main() {
	i = 0

	wg.Add(2)
	go MyFunc1()
	go MyFunc2()

	wg.Wait()
	fmt.Printf("final result: i = %d\n", i)
}
