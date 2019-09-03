package queue

import "sync"

type Item interface {
	// TODO Item 结构不对，需要重新定义
}

type ArrayQueue struct {
	item []Item
	lock sync.RWMutex
	len  int32
}

/**
	创建新的数组队列
 */
func (q *ArrayQueue) New() *ArrayQueue {
	q.item = []Item{}
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
func (q *ArrayQueue) Enqueue(E *Item) {
	q.lock.Lock()
	q.item = append(q.item, E )
	q.lock.Unlock()
}

/**
	出队
 */
func (q *ArrayQueue) Dequeue() Item {
	q.lock.Lock()
	item := q.item[q.len - 1]
	q.item = q.item[:len(q.item)-1 ]
	q.lock.Unlock()
	return &item
}
