package main_test

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	t.Cleanup(func() {
		fmt.Println("cleanup")
	})
	fmt.Println("testA")
}
