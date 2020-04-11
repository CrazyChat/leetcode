package structure

import "errors"

type Queue struct {
	cap int
	head int
	tail int
	data []int
}
// NewQueue 创建一个cap大小的队列
func NewQueue(capacity int) Queue {
	return Queue{cap: capacity, data: make([]int, capacity)}
}
// Enqueue 入队
func (q *Queue) Enqueue(value int) error {
	// (Tail + 1) % Cap == Head 表示队列满了
	if (q.tail+1) % q.cap == q.head {
		return errors.New("the queue is full")
	}
	q.data[q.tail] = value
	q.tail = (q.tail+1) % q.cap
	return nil
}
// Dequeue 出队
func (q *Queue) Dequeue() (int, error) {
	// 如果head == tail 表示队列为空
	if q.head == q.tail {
		return -1, errors.New("the queue is empty")
	}
	temp := q.data[q.head]
	q.head = (q.head + 1) % q.cap
	return temp, nil
}