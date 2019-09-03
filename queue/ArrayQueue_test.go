package queue

import "testing"

func TestArrayQueue_New(t *testing.T) {
	Q := new(ArrayQueue)
	if Q.len != 0 {
		t.Error("error")
	}
}

func TestArrayQueue_Enqueue(t *testing.T) {
	//Q := new(ArrayQueue)
	//Q.Enqueue()

}