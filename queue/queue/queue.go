package queue

import (
	"errors"
	"sync"
)

// Queue 队列
type Queue struct {
	data         []interface{} // 容器
	front        int           // 头
	rear         int           // 尾
	max          int           // 容器容量
	sync.RWMutex               // 读写锁
}

// InitQueue 初始化队列
func InitQueue(lenght int) (queue *Queue, err error) {
	if lenght < 1 {
		err = errors.New("长度不正确")
		return
	}
	queue = &Queue{
		data:  make([]interface{}, lenght+1),
		front: 0,
		rear:  0,
		max:   lenght + 1,
	}
	return
}

// Lenght 获取当前队列长度
func (q *Queue) Lenght() int {
	q.RLock()
	defer q.RUnlock()
	return (q.rear - q.front + q.max) % q.max
}

// IsFull 是否满
func (q *Queue) IsFull() bool {
	q.RLock()
	defer q.RUnlock()
	return q.isFull()
}

// isFull 是否满
func (q *Queue) isFull() bool {
	return ((q.rear + 1) % q.max) == q.front
}

// IsEmpty 是否空
func (q *Queue) IsEmpty() bool {
	q.RLock()
	defer q.RUnlock()
	return q.isEmpty()
}

// isEmpty 是否空
func (q *Queue) isEmpty() bool {
	return q.front == q.rear
}

// EnQueue 入队
func (q *Queue) EnQueue(item interface{}) error {
	q.Lock()
	defer q.Unlock()
	if q.isFull() {
		return errors.New("队列已满")
	}
	q.data[q.rear] = item
	q.rear = (q.rear + 1) % q.max
	return nil
}

// DeQueue 出队
func (q *Queue) DeQueue() (item interface{}, err error) {
	q.Lock()
	defer q.Unlock()
	if q.isEmpty() {
		err = errors.New("队列为空")
		return
	}
	item = q.data[q.front]
	q.front = (q.front + 1) % q.max
	return
}
