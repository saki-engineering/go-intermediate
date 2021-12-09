package main_test

import (
	"fmt"
	"testing"
)

func setup() {
	fmt.Println("setup")
}

func teardown() {
	fmt.Println("teardown")
}

func TestMain(m *testing.M) {
	setup()

	m.Run()

	teardown()
}

func TestA(t *testing.T) {
	fmt.Println("testA")
}

func TestB(t *testing.T) {
	fmt.Println("testB")
}
