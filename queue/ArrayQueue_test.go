package queue

import (
	"testing"
)

func TestArrayQueue_New(t *testing.T) {
	Q := ArrayQueue{}
	if Q.len != 0 {
		t.Error("error")
	}
}

func TestArrayQueue_Enqueue(t *testing.T) {
	Q := ArrayQueue{}
	var a int32 = 10
	Q.Enqueue(a)
	Q.Enqueue("Hello World")
	Q.Print()
}

func TestArrayQueue_Enqueue1(t *testing.T) {
	Q := ArrayQueue{}
	for i := 0; i < 20; i++ {
		Q.Enqueue(int32(i))
	}
	if Q.len != 20 {
		t.Error("队列长队错误")
	}
}

func TestArrayQueue_Dequeue(t *testing.T) {
	Q := ArrayQueue{}
	for i := 0; i < 20000; i++ {
		Q.Enqueue(int32(i))
	}
	var value interface{}
	for j := 0; j < 20000; j++ {
		value = Q.Dequeue()
		if value != int32(j) {
			t.Error("取出值不匹配")
		}
	}
}

func TestArrayQueue_Front(t *testing.T) {
	Q := ArrayQueue{}
	for i := 0; i < 20; i++ {
		Q.Enqueue(int32(i))
	}
	if Q.Front() != Q.Dequeue() {
		t.Error("Front方法错误")
	}
}
