package stack

import (
	"testing"
	"fmt"
)

func TestStack_Len(t *testing.T) {
	s := Stack{}
	fmt.Println("length: ", s.Len())
}

func TestStack_Push(t *testing.T) {
	s := Stack{}
	s.Push(2)
	if s.Len() != 1 {
		t.Error("count error")
	}
	s.Push("Hello World!")
	if s.Len() != 2 {
		t.Error("count error")
	}
	firstValue, err := s.Pop()
	if err != nil {
		t.Error(err)
	}
	if firstValue != "Hello World!"{
		t.Error("content error")
	}
	fmt.Println("firstValue: ", firstValue)
	fmt.Println("Push success!")
}