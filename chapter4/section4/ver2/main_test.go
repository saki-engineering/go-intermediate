package main_test

import (
	"fmt"
	"testing"
)

func TestTableDrivienParallelNG(t *testing.T) {
	defer func() {
		fmt.Println("cleanup")
	}()

	tests := []struct {
		testTitle string
	}{
		{testTitle: "subtest1"},
		{testTitle: "subtest2"},
		{testTitle: "subtest3"},
	}

	for _, test := range tests {
		testcase := test
		t.Run(testcase.testTitle, func(t *testing.T) {
			t.Parallel()
			fmt.Println(testcase.testTitle)
		})
	}
}

func TestTableDrivienParallelOK(t *testing.T) {
	t.Cleanup(func() {
		fmt.Println("cleanup")
	})

	tests := []struct {
		testTitle string
	}{
		{testTitle: "subtest1"},
		{testTitle: "subtest2"},
		{testTitle: "subtest3"},
	}

	for _, test := range tests {
		testcase := test
		t.Run(testcase.testTitle, func(t *testing.T) {
			t.Parallel()
			fmt.Println(testcase.testTitle)
		})
	}
}
