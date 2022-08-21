package main

import (
	"fmt"
	"strings"
)

func doubleInt(src int, intCh chan<- int) {
	result := src * 2
	intCh <- result
}

func doubleString(src string, strCh chan<- string) {
	result := strings.Repeat(src, 2)
	strCh <- result
}

func main() {
	ch1, ch2 := make(chan int), make(chan string)
	defer close(ch1)
	defer close(ch2)

	go doubleInt(1, ch1)
	go doubleString("hello", ch2)

	for i := 0; i < 2; i++ {
		select {
		case numResult := <-ch1:
			fmt.Println(numResult)
		case strResult := <-ch2:
			fmt.Println(strResult)
		}
	}
}
