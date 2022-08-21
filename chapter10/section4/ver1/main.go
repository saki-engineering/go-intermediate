package main

import "fmt"

func doubleInt(src int, intCh chan<- int) {
	result := src * 2
	intCh <- result
}

func main() {
	ch := make(chan int)
	defer close(ch)

	go doubleInt(1, ch)
	result := <-ch

	fmt.Println(result)
}
