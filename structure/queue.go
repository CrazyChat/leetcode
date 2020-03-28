package structure

import "errors"

type Queue struct {
	Cap int
	Head int
	Tail int
	DataSlice []interface{}
}
// NewQueue 创建一个cap大小的队列
func NewQueue(capacity int) Queue {
	return Queue{Cap: capacity, DataSlice: make([]interface{}, capacity)}
}
// Enqueue 入队
func (q *Queue) Enqueue(data interface{}) error {
	// (Tail + 1) % Cap == Head 表示队列满了
	if (q.Tail + 1) % q.Cap == q.Head {
		return errors.New("No Capcaity")
	}
	q.DataSlice[q.Tail] = data
	q.Tail = (q.Tail + 1) % q.Cap
	return nil
}
// Dequeue 出队
func (q *Queue) Dequeue() interface{} {
	// 如果head == tail 表示队列为空
	if q.Head == q.Tail {
		return nil
	}
	temp := q.DataSlice[q.Head]
	q.Head = (q.Head + 1) % q.Cap
	return temp
}