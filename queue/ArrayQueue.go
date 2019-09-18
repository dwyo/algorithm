package queue

import (
	"fmt"
	"sync"
)

type ArrayQueue struct {
	item []interface{}
	lock sync.RWMutex
	len  int32
}

/**
创建新的数组队列
*/
func (q *ArrayQueue) New() *ArrayQueue {
	q.len = 0
	return q
}

/**
判断队列是否为空
*/
func (q *ArrayQueue) IsEmpty() bool {
	return q.len == 0
}

/**
获取队列长度
*/
func (q *ArrayQueue) Size() int32 {
	return q.len
}

/**
入队
*/
func (q *ArrayQueue) Enqueue(e interface{}) {
	q.lock.Lock()
	q.item = append(q.item, e)
	q.len++
	q.lock.Unlock()
}

/**
出队
*/
func (q *ArrayQueue) Dequeue() interface{} {
	q.lock.Lock()
	item := q.item[0:1]
	q.item = q.item[1:q.len]
	q.len--
	q.lock.Unlock()
	return item[0]
}

/**
打印队列
 */
func (q *ArrayQueue) Print() {
	for idx, value := range q.item {
		fmt.Println(idx, value)
	}
}

/**
获取首位元素
 */
func (q *ArrayQueue) Front() interface{} {
	return q.item[0]
}
