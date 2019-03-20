package queue

import (
	"errors"
	"sync"
)

// 节点
type node struct {
	value interface{} // 值
	next  *node       // 下一个节点
}

// Queue 队列
type Queue struct {
	front      *node // 头
	rear       *node // 尾
	sync.Mutex       // 锁
}

// Enqueue 入队
func (q *Queue) Enqueue(item interface{}) error {
	q.Lock()
	defer q.Unlock()
	if item == nil {
		return errors.New("数据不可为空")
	}
	newNode := &node{value: item}
	if q.front == nil {
		q.front = newNode
	}
	if q.rear != nil {
		q.rear.next = newNode
	}
	q.rear = newNode
	return nil
}

// Dequeue 出队
func (q *Queue) Dequeue() (item interface{}, err error) {
	q.Lock()
	defer q.Unlock()
	if q.front == nil {
		err = errors.New("队列为空")
		return
	}
	item = q.front.value
	if q.front.next == nil {
		q.front = nil
	} else {
		q.front = q.front.next
	}
	return
}
