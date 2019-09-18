package linkedList

import (
	"github.com/davecgh/go-spew/spew"
)

// 节点
type Element struct {
	next, pre *Element
	value     interface{}
}

// 双向链表
type List struct {
	head *Element // 头
	tail *Element // 尾
	len  int
}

// @Summary 向链表添加一个元素
func (l *List) Append(e *Element) bool {
	if e == nil {
		return false
	}
	e.next = nil

	if l.len == 0 {
		e.pre = nil
		l.head = e
	} else {
		e.pre = l.tail
		l.tail.next = e
	}
	l.tail = e
	l.len++
	return true
}

// @Summary 在某节点后添加一个节点
func (l *List) AppendNode(e *Element, idx int) bool {
	if l.len < idx {
		return false
	}
	cur := l.head
	for i := 0; i < idx; i++ {
		cur = cur.next
	}

	if cur.next != nil {
		e.next = cur.next
		cur.next.pre = e
	}
	e.pre = cur
	cur.next = e
	l.len++

	return true
}

// @Summary 从头开始遍历节点
func (l *List) Each()  {
	if l.len == 0 {
		return
	}
	cur := l.head
	for cur.next != nil {
		spew.Dump(cur.value)
		cur = cur.next
	}
}

// @Summary 从尾开始遍历节点
func (l *List) Reversion() {
	if l.len == 0 {
		return
	}
	cur := l.tail
	for cur.pre != nil {
		spew.Dump(cur.value)
		cur = cur.pre
	}
}

// @Summary 删除一个节点
func (l *List)Del(idx int) bool {
	if l.len < idx {
		return false
	}
	cur := l.head
	for i:= 0; i < idx - 1; i++{
		cur = cur.next
	}
	cur.next.next.pre = cur
	cur.next = cur.next.next
	l.len--
	return true
}

// @Summary 获取链表长度
func (l *List) GetLen() int {
	return l.len
}

// @Summary 判断链表是否为空
func (l *List) IsEmpty() bool {
	return l.len == 0
}


